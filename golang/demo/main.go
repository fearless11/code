package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hi", hiFunc)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hiFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
