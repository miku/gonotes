package main

import (
	"errors"
	"log"
)

var ErrTooSmall = errors.New("too small")
var ErrTooLarge = errors.New("too large")

func buggy(i int) error {
	if i < 10 {
		return ErrTooSmall
	} else if i > 20 {
		return ErrTooLarge
	}
	return nil
}

func main() {
	var i = 100
	err := buggy(i)
	switch {
	case err == ErrTooSmall:
		log.Println("small")
	case err == ErrTooLarge:
		log.Println("large")
	default:
		log.Fatal("unexpected error")
	}
}
