package main

import "fmt"

func main() {
	a := [3]int{0, 1, 2}
	var b [3]int
	s := a[:]
	fmt.Printf("%T %T\n", a, s)

	copy(b[:], s)
	fmt.Println(b)
}
