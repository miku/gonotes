package main

import (
	"fmt"
)

func sliceInfo(name string, s []byte) {
	fmt.Printf("[%s] %p len=%d cap=%d\n", name, s, len(s), cap(s))
}

func main() {
	a := []byte("Help")
	b := append(a, []byte(" Me")...)
	c := append(a, []byte(" Him")...)
	fmt.Println(string(a), "-", string(b), "-", string(c))
}

// func main() {
// 	a := []byte("Help")
// 	sliceInfo("a", a)
//
// 	b := append(a, []byte(" Me")...)
// 	sliceInfo("a", a)
// 	sliceInfo("b", b)
//
// 	c := append(a, []byte(" Him")...)
// 	sliceInfo("a", a)
// 	sliceInfo("b", b)
// 	sliceInfo("c", c)
//
// 	fmt.Println(string(a), "-", string(b), "-", string(c))
// }
