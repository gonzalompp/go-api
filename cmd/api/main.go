package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// Env var key for the message content
const EnvHelloMessage = "SECRET_HELLO_MESSAGE"

func HelloServer(w http.ResponseWriter, r *http.Request) {
	// values
	name := r.URL.Path[1:]
	message := os.Getenv(EnvHelloMessage)

	// response
	fmt.Fprintf(w, message, name)
}
