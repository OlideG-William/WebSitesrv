package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", helloHandleFunc)
	http.HandleFunc("/about", aboutHandleFunc)

	http.ListenAndServe(":8080", nil)
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "r.URL.Path:, %s", r.URL.Path)
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our Gopher powered website.")
}
