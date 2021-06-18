package main

import (
	"log"
	"net/http"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!\n"))
}

func main() {
	log.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", &handler{}))
}
