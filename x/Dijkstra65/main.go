// Solution of a Problem in Concurrent Program Control,
// http://rust-class.org/static/classes/class19/dijkstra.pdf,
// https://asciinema.org/a/GX0ntCTr5hFf6CzHDvs6VaeU6
package main

import (
	"flag"
	"time"
)

var (
	N    = flag.Int("n", 3, "number of processes")
	T    = flag.Duration("t", 1000*time.Millisecond, "exit simulation after t (e.g. 1s, 100ms, ...)")
	Cdur = flag.Duration("C", 10*time.Millisecond, "critical section duration")
)

func main() {
	flag.Parse()
	var (
		k        int = 0
		b            = boolSlice(*N, true)
		c            = boolSlice(*N, true)
		computer     = func(i int) {
		Li0:
			b[i] = false
		Li1:
			if k != i {
				c[i] = true
				if b[k] {
					k = i
				}
				goto Li1
			} else {
				c[i] = false
				for j := 0; j < *N; j++ {
					if j != i && !c[j] {
						goto Li1
					}
				}
			}
			time.Sleep(*Cdur) // simulate critical section
			c[i] = true
			b[i] = true
			goto Li0
		}
	)
	for i := 0; i < *N; i++ {
		go computer(i)
	}
	time.Sleep(*T)
}

func boolSlice(n int, iv bool) []bool {
	b := make([]bool, n)
	for i := 0; i < len(b); i++ {
		b[i] = iv
	}
	return b
}
