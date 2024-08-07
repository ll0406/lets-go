package main

import (
	"log"
	"net/http"
)

// Define home handler writes a byte slice
// "Hello from Snippetbox" as res
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet handler function
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
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
