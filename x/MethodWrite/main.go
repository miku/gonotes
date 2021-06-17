package main

import (
	"fmt"
	"log"
)

type ByteSlice []byte

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	*p = data
	return len(data), nil
}

func (p *ByteSlice) String() string {
	return string(*p)
}

func main() {
	var bs ByteSlice

	_, err := fmt.Fprintf(&bs, "This hour has %d days", 7)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(&bs)
}
