package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("starting server...")
	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	fmt.Print(dir)
	router := mux.NewRouter()
	// 1. HTMX-based Request
	router.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, "<button id=\"button\" class=\"doit\" hx-post=\"/clickedagain\" hx-swap=\"outerHTML\" >CLICK IT!!!</button>")
	}).Methods("POST")

	// 2. HTMX-based Request
	router.HandleFunc("/clickedagain", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, "<button id=\"button\" class=\"doit yellow\">You did it!</button>")
	}).Methods("POST")

	// Serving the static assets
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	http.ListenAndServe(":8080", router)
}
