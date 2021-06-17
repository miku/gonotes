package main

import (
	"fmt"
)

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	y := append(x, 2)
	z := append(x, 3)
	fmt.Println(y, z) // ?? [0 1 2] [0 1 3]
}

func b() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y, z) // [0 1 2 3 4] [0 1 2 3 4] // [0 1 2 3], [0 1 2 3 4] // [0 1 2 4] [0 1 2 4] (*)
}

func main() {
	a()
	B()
}

func B() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)

	sliceInfo("x", x)
	y := append(x, 3)

	sliceInfo("x", x)
	sliceInfo("y", y)

	z := append(x, 4)
	sliceInfo("x", x)
	sliceInfo("y", y)
	sliceInfo("z", z)

	fmt.Println(y, z)
}
func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d\n", name, s, len(s), cap(s))
}

// [x] 0xc00001c200 len=3 cap=4
// [x] 0xc00001c200 len=3 cap=4

// [y] 0xc00001c200 len=4 cap=4

// [x] 0xc00001c200 len=3 cap=4
// [y] 0xc00001c200 len=4 cap=4
// [z] 0xc00001c200 len=4 cap=4
// [0 1 2 4] [0 1 2 4]
