package main

import "time"

func a() int {
	time.Sleep(10 * time.Millisecond)
	return 42
}

func main() {}
