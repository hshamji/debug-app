package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	b := new(bytes.Buffer)
	for key, value := range req.Header {

		fmt.Println(key, value)
		fmt.Fprintf(b, "%s=%s", key, value)
	}

	w.Write([]byte(b.Bytes()))
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health Check</h1>")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":10002", nil)
}
