package main

import "fmt"

func main() {
	var x *ErrorX = nil
	var err error = x
	if err == nil {
		fmt.Println("nope, this won't happen ")
	} else {
		fmt.Printf("here's my non-nil error: %#v %T\n", x, x)
	}
}

type ErrorX struct{}

func (x *ErrorX) Error() string {
	return "hi i'm ErrorX"
}
