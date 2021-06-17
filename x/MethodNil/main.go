package main

import "fmt"

// IntList is a linked list. A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

func (l *IntList) Sum() int {
	if l == nil {
		return 0
	}
	return l.Value + l.Tail.Sum()
}

func main() {
	list := &IntList{1, &IntList{2, nil}}
	fmt.Println(list.Sum())
}
