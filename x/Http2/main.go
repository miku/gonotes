package main

// Making a request.

import (
	"crypto/sha1"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	u := "https://heise.de"
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// b, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(b))

	h := sha1.New()
	r := io.TeeReader(resp.Body, h)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	log.Printf("sha1: %x", h.Sum(nil))
}
