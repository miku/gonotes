package main

import (
	"fmt"
	"net/http"
)

// Exercise: Write a debug handler, that returns a response with information
// about the request. You can use the helper function: httputil.DumpRequest(r,
// false) - or return selected fields.
//
// Return an HTTP 500, if something goes wrong.

// Debug show the request.
func Debug(w http.ResponseWriter, r *http.Request) {
}

// Echo is a basic HTTP Handler.
func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}

func main() {
	// Register Handler (e.g. with the default servemux), start a server
}
