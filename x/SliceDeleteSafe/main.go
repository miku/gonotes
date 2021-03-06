package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a []*string
	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		a = append(a, &s)
	}
	i := 2
	dump(a)
	copy(a[i:], a[i+1:])
	a[len(a)-1] = nil // or the zero value of T
	a = a[:len(a)-1]
	dump(a)
}

func dump(ss []*string) {
	fmt.Printf("[")
	for _, s := range ss {
		fmt.Printf("%s", *s)
	}
	fmt.Printf("]")
	fmt.Println()
}
