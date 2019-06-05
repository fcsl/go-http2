package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, HTTP2!")
}

func main() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", handler)
	server := http.Server{
		Addr:    ":8443",
		Handler: serverMux,
	}
	server.ListenAndServeTLS("server.pem", "server.key")
}
