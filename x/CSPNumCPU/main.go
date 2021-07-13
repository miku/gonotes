package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// NumCPU returns the number of logical
	// CPUs usable by the current process.
	fmt.Println(runtime.NumCPU())

	go func() {
		fmt.Println("started")
		time.Sleep(10 * time.Millisecond)
	}()

	fmt.Println(runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
	fmt.Println(runtime.NumGoroutine())

}
