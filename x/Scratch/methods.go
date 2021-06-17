package main

import (
	"fmt"
	"sync"
)

type Point struct {
	sync.Mutex // embedding

	X int
	Y int

	// mu    sync.Mutex // over the protected field
	state int
}

func (p *Point) Scale(factor int) {
	p.Lock()
	defer p.Unlock()
	p.state = 1
	p.X = p.X * factor
	p.Y = p.Y * factor
	return p
}

// type MyInt int
//
// type transform func(string) string
//
// func (m MyInt) IsEven() bool {
// 	return m%2 == 0
// }
//
// func identity(s string) string {
// 	return s
// }

func main() {
	p := Point{X: 1, Y: 1}
	fmt.Printf("%+v\n", p)
	p = p.Scale(10)
	fmt.Printf("%+v\n", p)

	// var f transform = identity
	// fmt.Printf("%T %v\n", f, f)
	// v := MyInt(124)
	// fmt.Printf("%v %v\n", v, v.IsEven())
}

// Ex:
//
// data
// |
// v
// []byte -> []string // extract some fields --- F
//
//           export to TSV
//
