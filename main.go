package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("hello working"))
	})
	log.Println("Server started")
	http.ListenAndServe(":80", nil)
}
