package main

import "fmt"

type servers struct {
	i int
	s []string
}

func (s *servers) next() string {
	s.i = (s.i + 1) % len(s.s)
	return s.s[s.i]
}

func Select(f func() string) {
	fmt.Println(f())
}

func fixed() string {
	return "fixed"
}

func main() {
	s := servers{s: []string{"a", "b", "c"}}
	Select(s.next)
	Select(s.next)
	Select(s.next)
	Select(s.next)
	Select(s.next)
	Select(s.next)
	Select(fixed)
	Select(fixed)
}
