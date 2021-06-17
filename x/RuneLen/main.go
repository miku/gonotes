package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Schöneberg" // https://www.utf8-zeichentabelle.de/unicode-utf8-table.pl?names=-&utf8=dec&unicodeinhtml=dec

	fmt.Printf("%s -- %d\n", s, len(s))                    // 11
	fmt.Printf("%s -- %d\n", s, len([]rune(s)))            // 10
	fmt.Println("rune count: ", utf8.RuneCountInString(s)) // 10
	fmt.Println("b", []byte(s))
	fmt.Println("r", []rune(s))

}
