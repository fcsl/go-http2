// +build ignore

package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

type serverHandler struct {
}

func (sh *serverHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	w.Header().Set("server", "h2test")
	w.Write([]byte("http2 without tls"))
}

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      &serverHandler{},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	s2 := &http2.Server{
		IdleTimeout: 1 * time.Minute,
	}
	http2.ConfigureServer(server, s2)
	l, _ := net.Listen("tcp", ":8080")
	defer l.Close()
	fmt.Println("Start server...")
	for {
		rwc, err := l.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		go s2.ServeConn(rwc, &http2.ServeConnOpts{BaseConfig: server})

	}
}
