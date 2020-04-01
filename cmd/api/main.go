package main

import (
	"fmt"
	"net/http"
	"os"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {

	logctx := log.WithFields(log.Fields{
		"application": "go-api",
		"app": "api",
	})

	logctx.Info("starting application")

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// Env var key for the message content
const EnvHelloMessage = "SECRET_HELLO_MESSAGE"

func HelloServer(w http.ResponseWriter, r *http.Request) {

	logctx := log.WithFields(log.Fields{
		"application": "go-api",
		"app": "api",
		"request_url": r.URL,
		"request_host": r.Host,
		"request_method": r.Method,
		"request_user_agent": r.UserAgent(),
	})

	logctx.Info("request received")
	// values
	name := r.URL.Path[1:]
	message := os.Getenv(EnvHelloMessage)

	logctx.Infof("user name %s",name)

	// response
	fmt.Fprintf(w, message, name)
	logctx.Info("request done")
}
