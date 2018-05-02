package server

import (
	"io/ioutil"
	"net/http"
	"github.com/apex/log"
	"io"
	"time"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"bytes"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"strconv"
	"github.com/satori/go.uuid"
	"strings"
	"path/filepath"
	"os"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"context"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type Env struct {
	S3Region string
	S3Bucket string
	S3Prefix string
	S3TrailingPrefix string
	S3Timeout time.Duration
	S3Session *session.Session
	S3Svc s3iface.S3API
}

func Get(w http.ResponseWriter, r *http.Request) {
	defer dclose(r.Body)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "HTTP endpoint for S3: %s\n", r.URL.Path)
}

func (env *Env) Post(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer dclose(r.Body)

	successful := Upload(env, b)

	if successful {
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func dclose(c io.Closer) {
	if c == nil {
		return
	}
	if err := c.Close(); err != nil {
		log.WithError(err).Error("unable to close response")
	}
}

// collects log messages from fluent-bit and sends them to S3
func Upload(env *Env, body []byte) bool {
	//renew expired or missing session
	if env.S3Session == nil || env.S3Session.Config.Credentials.IsExpired() {
		env.S3Session = GetSession(env.S3Region)
	}
	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	env.S3Svc = s3.New(env.S3Session)

	t := time.Now()
	objectKey := GenerateObjectKey(env.S3Prefix, env.S3TrailingPrefix, t)

	// Upload
	return PutLogBlobToS3(env, body, objectKey)
}

// format is S3_PREFIX/S3_TRAILING_PREFIX/date/hour/timestamp_uuid.log
func GenerateObjectKey(S3Prefix, S3TrailingPrefix string, t time.Time) string {
	timestamp := t.Format("20060102150405")
	date := t.Format("20060102")
	hour := strconv.Itoa(t.Hour())
	logUUID := uuid.Must(uuid.NewV4()).String()
	fileName := strings.Join([]string{timestamp, "_", logUUID, ".log"}, "")

	objectKey := filepath.Join(S3Prefix, S3TrailingPrefix, date, hour, fileName)
	return objectKey
}

func GetSession(S3Region string) *session.Session {

	log.Error("Should not be here!")
	metasession, err := session.NewSession(&aws.Config{Region: aws.String(S3Region)})
	if err != nil {
		log.WithError(err).Fatal("error creating S3 metasession")
		os.Exit(1)
	}
	metadata := ec2metadata.New(metasession)
	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.EnvProvider{},
			&credentials.SharedCredentialsProvider{},
			&ec2rolecreds.EC2RoleProvider{Client: metadata},
		})
	awsConfig := aws.NewConfig()
	awsConfig.WithCredentials(creds)
	awsConfig.WithRegion(S3Region)
	sess, err := session.NewSession(awsConfig)
	// Create a single AWS session (we can re use this if we're uploading many files)
	if err != nil {
		log.WithError(err).Fatal("error creating S3 session")
		os.Exit(1)
	}

	return sess
}


// Write blob to S3://S3_BUCKET/objectKey with timeout of S3_TIMEOUT
func PutLogBlobToS3(env *Env, body []byte, objectKey string) bool {
	// Create a context with a timeout that will abort the upload if not
	// completed by the time specified in timeout.
	ctx := context.Background()
	var cancelFn func()
	if env.S3Timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, env.S3Timeout)
		// Ensure the context is canceled to prevent leaking.
		// See context package for more information, https://golang.org/pkg/context/
		defer cancelFn()
	}


	// Uploads the object to S3. The Context will interrupt the request if the
	// timeout expires.
	_, err := env.S3Svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(env.S3Bucket),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(body),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			log.WithError(err).Errorf("upload canceled due to timeout to s3://%s/%s", env.S3Bucket, objectKey)
		} else {
			log.WithError(err).Errorf("error uploading to s3://%s/%s", env.S3Bucket, objectKey)
		}
	} else {
		log.Infof("Successfully wrote s3://%s/%s", env.S3Bucket, objectKey)

	}
	return err == nil
}



