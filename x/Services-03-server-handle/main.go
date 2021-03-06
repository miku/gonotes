package main

import (
	"fmt"
	"log"
	"net/http"
)

// Echo is a basic HTTP Handler.
func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}

func main() {

	log.Println("http://localhost:8000")

	// Handle registers the handler for the given pattern in the DefaultServeMux.
	http.Handle("/index", http.HandlerFunc(Echo))

	// Start a server listening on port 8000 and responding using Echo.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
