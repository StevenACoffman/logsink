package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/gorilla/pat"
	"time"
	"strconv"
)


func init() {
	log.SetHandler(json.Default)
}

var delay int

func main() {
	port := getEnv("PORT", "3000")

	if value, ok := os.LookupEnv("DELAY"); ok {
		if delayResponse, err := strconv.Atoi(value); err == nil {
			delay = delayResponse
		}
	}

	addr := ":" + port
	app := pat.New()
	app.Post("/", post)
	log.WithField("status", "Starting Listener on PORT: "+port).Info("status")
	if err := http.ListenAndServe(addr, app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func post(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	time.Sleep(time.Duration(delay) * time.Second)

	if string(b) == "" {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusBadRequest)
		return
	}


	return
}
