package main

import "fmt"

func main() {

	var x = []int{1, 2, 3}
	var y []int
	var z = make([]int, 10)

	fmt.Println(x, y, z)

	y = append(y, 123)
	y = append(y, 234)
	y = append(y, 345)

	fmt.Println(x, y, z)

	z = append(z, 1)
	fmt.Println(x, y, z)

	fmt.Println(z[0])
	fmt.Println(x[1:])

}
