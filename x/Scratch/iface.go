package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type MyReader interface {
	Read([]byte) (int, error)
}

type zeroReader struct{}

func (z zeroReader) Read(p []byte) (int, error) {
	return 0, io.EOF
}

func main() {
	f, err := os.Open("gc.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T %#v\n", f, f)

	var r MyReader
	fmt.Printf("%T %#v\n", r, r)

	r = f
	// b, _ := ioutil.ReadAll(r)
	// fmt.Println(string(b))

	z := zeroReader{}
	r = z

	fmt.Printf("%T %#v\n", r, r)
}
