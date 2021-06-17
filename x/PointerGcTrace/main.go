package main

import "fmt"

func main() {
	var v map[string]string
	for i := 0; i < 1000; i++ {
		v = make(map[string]string)
		fmt.Println(v)
	}
}

// No of runs, time since start, time spent in GC, timings, heaps, threads used

// $ GODEBUG=gctrace=1 go run main.go > /dev/null
// gc 1 @0.005s 5%: 0.077+0.98+0.13 ms clock, 0.62+0.47/0.48/0.13+1.0 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
// gc 2 @0.010s 3%: 0.017+0.27+0.003 ms clock, 0.13+0.19/0.27/0.45+0.026 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 3 @0.020s 2%: 0.022+0.30+0.011 ms clock, 0.17+0.12/0.36/0.65+0.092 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 4 @0.029s 2%: 0.029+0.36+0.008 ms clock, 0.23+0.19/0.36/0.50+0.069 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 5 @0.037s 1%: 0.020+0.33+0.008 ms clock, 0.16+0.12/0.41/0.73+0.066 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 6 @0.045s 1%: 0.027+0.66+0.009 ms clock, 0.21+0.16/0.55/0.82+0.074 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
// gc 7 @0.047s 2%: 0.030+0.29+0.010 ms clock, 0.24+0.33/0.42/0.74+0.082 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 8 @0.052s 1%: 0.017+0.40+0.002 ms clock, 0.14+0.096/0.50/0.87+0.018 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 9 @0.055s 2%: 0.036+0.60+0.084 ms clock, 0.28+0.22/0.62/0.62+0.67 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
// gc 10 @0.057s 2%: 0.020+0.34+0.002 ms clock, 0.16+0.23/0.51/0.91+0.018 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// # command-line-arguments
// gc 1 @0.002s 4%: 0.009+0.94+0.002 ms clock, 0.073+0.24/1.1/0.79+0.016 ms cpu, 4->5->4 MB, 5 MB goal, 8 P
// # command-line-arguments
// gc 1 @0.001s 9%: 0.039+0.81+0.002 ms clock, 0.31+0.082/1.1/0.12+0.022 ms cpu, 4->6->5 MB, 5 MB goal, 8 P
// gc 2 @0.006s 5%: 0.011+0.77+0.008 ms clock, 0.091+0.091/1.2/0.95+0.067 ms cpu, 8->9->8 MB, 10 MB goal, 8 P
