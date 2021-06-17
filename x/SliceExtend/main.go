package main

import "fmt"

func main() {

	var a = []int{0, 1, 2, 3}
	sliceInfo("a", a)

	a = append(a, make([]int, 10)...)

	sliceInfo("a", a)

	a[13] = 13
	fmt.Printf("%v\n", a)
}

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d\n", name, s, len(s), cap(s))
}
