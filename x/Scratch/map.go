package main

import "fmt"

func main() {

	var m = make(map[string]string)

	m["a1"] = "def"
	m["b2"] = "xyz"
	m["c3"] = "ijk"

	for k, v := range m {
		fmt.Printf("%s => %s\n", k, v)
	}

	for k := range m {
		fmt.Printf("%s\n", k)
	}

	// all keys as slice
	var keys []string
	for k := range m {
		keys = append(keys, k) //
	}
	fmt.Printf("%v\n", keys)
}
