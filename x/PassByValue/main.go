package main

import "fmt"

func f(v int) {
	v += 1
}

func g(v *int) {
	*v += 1
}

func main() {
	v := 1
	f(v)
	fmt.Println(v)
	g(&v)
	fmt.Println(v)
}
