package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define home handler writes a byte slice
// "Hello from Snippetbox" as res
func home(w http.ResponseWriter, r *http.Request) {
	// Checks the URL and render 404 if needed
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet handler function
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Restrict createSnippet to take only POST request
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// WriteHeader and Write can only be called once
		// Body and Header modifications need to happen before it
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Create a snippet..."))
}

func main() {
	mux := http.NewServeMux()
	// Acting like /** (Subtree path ends with / catch all)
	mux.HandleFunc("/", home)
	// Fixed path
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
