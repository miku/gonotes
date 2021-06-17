package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// 1 map init
	m := make(map[string]int)

	// 2 iteration over words

	for _, w := range strings.Fields(s) {
		m[w]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
