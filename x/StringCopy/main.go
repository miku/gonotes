package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	v := make([]byte, 1_073_741_824)
	s := string(v) // Is this a copy? Yes.
	fmt.Printf("allocated []byte, string of len %v %v\n", len(s), time.Since(t))
}
