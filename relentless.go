package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.Handle("/", http.FileServer(http.Dir("app")))
	http.ListenAndServe(":8080", nil)
}

type signup struct {
	Name    string        `json:"name"`
	Email   string        `json:"email"`
	Options signupOptions `json:"options"`
}

type signupOptions struct {
	Gap bool
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.Form.Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}
