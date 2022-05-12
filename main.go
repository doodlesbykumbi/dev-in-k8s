package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hell thenox\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	port := "8080"
	if val, ok := os.LookupEnv("PORT"); ok {
		port = val
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatal(err)
	}

	// Log server started
	log.Printf("Server started at port %v", listener.Addr().(*net.TCPAddr).Port)
	panic(http.Serve(listener, nil))
}
