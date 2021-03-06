package main

import "fmt"

func main() {

	var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sliceInfo("a", a)

	keep := func(x int) bool {
		return x%2 == 0
	}

	n := 0
	for _, x := range a { // i declared but not used
		if keep(x) {
			a[n] = x
			n++
		}
	}
	sliceInfo("a", a)
	a = a[:n]
	sliceInfo("a", a)
}

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d %v\n", name, s, len(s), cap(s), s)
}

// [a] 0xc0000220a0 len=10 cap=10 [0 1 2 3 4 5 6 7 8 9]
// [a] 0xc0000220a0 len=10 cap=10 [0 2 4 6 8 5 6 7 8 9]
// [a] 0xc0000220a0 len=5 cap=10 [0 2 4 6 8]
