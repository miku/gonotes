package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// a = append(a[:i], a[j:]...)

	a = append(a[:2], a[7:]...) // [a] 0xc0000be000 len=4 cap=9 [1 2 8 9]

	// fmt.Println(a)
	sliceInfo("a", a)
	a = a[:5]

	sliceInfo("a", a)
	// a = a[:50] // panic: runtime error: slice bounds out of range [:50] with capacity 9
	// sliceInfo("a", a)
	// fmt.Println(a)
}

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d %v\n", name, s, len(s), cap(s), s)
}
