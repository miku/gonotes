package main

import "fmt"

func main() {
	var v map[string]string
	for i := 0; i < 1000; i++ {
		v = make(map[string]string)
		fmt.Println(v)
	}
}
