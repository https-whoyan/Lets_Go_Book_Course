package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from home!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetView!"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetCreate!"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet/view", snippetView)
	http.HandleFunc("/snippet/create", snippetCreate)
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}
