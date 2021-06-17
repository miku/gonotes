package main

import "fmt"

func gen() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func square(c chan int) chan int {
	ch := make(chan int)
	go func() {
		for v := range c {
			ch <- v * v
		}
		close(ch)
	}()
	return ch
}

func main() {
	for i := range square(gen()) {
		fmt.Println(i)
	}
}
