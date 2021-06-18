package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("Hello World!\n"))
		}),
	))
}
