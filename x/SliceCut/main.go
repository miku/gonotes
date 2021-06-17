package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// a = append(a[:i], a[j:]...)
	a = append(a[:2], a[7:]...)
	// fmt.Println(
	// sliceInfo("a", a)
	// a = a[:5]
	// sliceInfo("a", a)
	// fmt.Println(a)
	// a = a[:50]
	// sliceInfo("a", a)
	// fmt.Println(a)
}

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d\n", name, s, len(s), cap(s))
}
