package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"goji.io"
	"goji.io/pat"
)

func main() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/hello/:name"), hello)
	http.Handle("/", http.FileServer(http.Dir("app")))

	http.ListenAndServe("localhost:8000", mux)
}

type signup struct {
	Name    string        `json:"name"`
	Email   string        `json:"email"`
	Options signupOptions `json:"options"`
}

type signupOptions struct {
	Gap bool `json:"gap"`
}

func hello(c context.Context, w http.ResponseWriter, r *http.Request) {
	name := r.Form.Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}
