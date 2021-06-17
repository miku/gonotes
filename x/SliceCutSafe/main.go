package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var a []*string
	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		a = append(a, &s)
	}
	i, j := 2, 7
	dump(a)

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k++ {
		log.Printf("setting s[%d] to nil", k)
		a[k] = nil // or the zero value of T
	}
	a = a[:len(a)-j+i]
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
