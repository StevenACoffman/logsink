package main

import (
	"fmt"
	"github.com/StevenACoffman/logsink/go-s3-logsink/server"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/gorilla/pat"
	"net/http"
	"os"
	"strconv"
	"time"
)

func init() {
	log.SetHandler(json.Default)
}

func main() {
	env := &server.Env{
		S3Region:         getRequiredEnv("AWS_REGION"),
		S3Bucket:         getRequiredEnv("S3_BUCKET"),
		S3Prefix:         getRequiredEnv("S3_PREFIX"),        //S3 Prefix does not contain the protocol
		S3TrailingPrefix: os.Getenv("S3_TRAILING_PREFIX"),    //Optional
		S3Timeout:        getDurationEnv("S3_TIMEOUT", "60"), //Optional
	}
	env.S3Session = server.GetSession(env.S3Region)

	port := getEnv("PORT", "3000")
	addr := ":" + port
	app := pat.New()
	app.Get("/", server.Get)
	app.Post("/", env.Post)
	log.WithField("status", "Starting Listener on PORT: "+port).Info("status")
	if err := http.ListenAndServe(addr, app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}

func getDurationEnv(key, fallback string) time.Duration {
	value := getEnv(key, fallback)
	if intValue, err := strconv.Atoi(value); err == nil {
		return time.Duration(intValue) * time.Second
	}
	log.Warnf("Unable to parse duration from environment variable %s with value %s", key, value)
	return time.Duration(60) * time.Second
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getRequiredEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Fatalf("Required Environment variable %s unset", key)
	panic(fmt.Sprintf("Required Environment variable %s unset", key))
}
