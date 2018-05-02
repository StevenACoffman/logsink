package main_test

import (
	"github.com/StevenACoffman/logsink/go-s3-logsink/server"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/gorilla/pat"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHello(t *testing.T) {

	app := pat.New()
	app.Get("/", server.Get)
	s := httptest.NewServer(app) // http provides its own test objects
	defer s.Close()

	resp, err := http.Get(s.URL) // test from the outside
	if err != nil {
		t.Error(err)
	}
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Error(err)
	} else if string(body) != "HTTP endpoint for S3: /\n" {
		t.Error("expected", "HTTP endpoint for S3: /", "got", string(body))
	}
}

// Define a mock struct to be used in your unit tests of myFunc.
type mockS3Client struct {
	s3iface.S3API
}

func (m *mockS3Client) PutObjectWithContext(aws.Context, *s3.PutObjectInput, ...request.Option) (*s3.PutObjectOutput, error) {
	return &s3.PutObjectOutput{}, nil
}

func TestUpload(t *testing.T) {
	//mockserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Print("Got here!")
	//	w.WriteHeader(http.StatusOK)
	//}))
	//
	//var sess = session.Must(session.NewSession(aws.NewConfig().
	//	WithCredentials(credentials.NewStaticCredentials("AKID", "SECRET", "SESSION")).
	//	WithRegion("mock-region").
	//	WithEndpoint(mockserver.URL).
	//	WithHTTPClient(mockserver.Client()).
	//	WithDisableSSL(true)))

	var env = &server.Env{
		S3Region:         "notaregion",
		S3Bucket:         "notabucket",
		S3Prefix:         "notaprefix", //S3 Prefix does not contain the protocol
		S3TrailingPrefix: "trailing",   //Optional
		S3Svc:            &mockS3Client{},
	}

	current := time.Now()
	objectKey := server.GenerateObjectKey(env.S3Prefix, env.S3TrailingPrefix, current)

	payload := "{}"

	successful := server.PutLogBlobToS3(env, []byte(payload), objectKey)
	if !successful {
		t.Error("Failed to upload to mock server")
	}

}
