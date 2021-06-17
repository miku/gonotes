package main

import "log"

func hello() string {
	return "hello"
}

func hi() string {
	return "hi"
}

func main() {
	var greeters = map[string]func() string{
		"hello": hello,
		"hi":    hi,
	}

	greeter := "hi"
	f := greeters[greeter]
	log.Println(f())
}
