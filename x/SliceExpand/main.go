package main

import "fmt"

func main() {
	var a = []int{0, 1, 2, 3}
	sliceInfo("a", a)
	i := 3
	j := 4
	//  --------------------------------------------------
	//                --------------------------------
	a = append(a[:i], append(make([]int, j), a[i:]...)...)
	sliceInfo("a", a)
	fmt.Println(a)
}

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d\n", name, s, len(s), cap(s))
}
