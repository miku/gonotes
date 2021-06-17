package main

import (
	"fmt"
	"log"
)

type Greeter interface {
	Greet() string
}

type English struct{}

func (e English) Greet() string {
	return "Hello"
}

type German struct{}

func (g German) Greet() string {
	return "Guten Tag"
}

func main() {
	var g Greeter = German{}

	// Type assertion
	concrete, ok := g.(German) // panic: interface conversion: main.Greeter
	fmt.Printf("%T %v\n", concrete, ok)

	if _, ok := g.(German); !ok {
		fmt.Println("conversion failed")
	}

	switch t := g.(type) {
	case English:
		log.Printf("%T %t", t, t)
	case German:
		log.Printf("%T %t", t, t)
	default:
		log.Println("unknown type")
	}
}
