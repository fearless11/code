package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type hi struct{}

func (h hi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi world")
}

func main() {
	hiHandler := hi{}
	http.Handle("/hi", hiHandler)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
