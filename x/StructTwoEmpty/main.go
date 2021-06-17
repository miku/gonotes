package main

import "fmt"

func main() {
	var a, b struct{}
	fmt.Println(&a == &b) // false, but may be true?
}
