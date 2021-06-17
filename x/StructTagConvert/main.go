package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string `json:"name"`
}

func main() {
	a := A{}
	b := B{}

	fmt.Printf("%T %T", a, b)

	c := A(b)
	fmt.Println(c)
}
