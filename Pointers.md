# Pointers

Pointers provide a way to share data across program boundaries. Having the ability to share and reference data with a pointer provides the benefit of efficiency. There is only one copy of the data and everyone can see it changing. The cost is that anyone can change the data which can cause side effects in running programs.

## Notes

* Use pointers to share data.
* Values in Go are always pass by value.
* "Value of", what's in the box. "Address of" ( **&** ), where is the box.
* The (*) operator declares a pointer variable and the "Value that the pointer points to".

## Escape Analysis

* When a value could be referenced after the function that constructs the value returns.
* When the compiler determines a value is too large to fit on the stack.
* When the compiler doesn’t know the size of a value at compile time.
* When a value is decoupled through the use of function or interface values.

## Garbage Collection History

The design of the Go GC has changed over the years:
* Go 1.0, Stop the world mark sweep collector based heavily on tcmalloc.
* Go 1.2, Precise collector, wouldn't mistake big numbers (or big strings of text) for pointers.
* Go 1.3, Fully precise tracking of all stack values.
* Go 1.4, Mark and sweep now parallel, but still stop the world.
* Go 1.5, New GC design, focusing on latency over throughput.
* Go 1.6, GC improvements, handling larger heaps with lower latency.
* Go 1.7, GC improvements, handling larger number of idle goroutines, substantial stack size fluctuation, or large package-level variables.
* Go 1.8, GC improvements, collection pauses should be significantly shorter than they were in Go 1.7, usually under 100 microseconds and often as low as 10 microseconds.
* Go 1.9, Large object allocation performance is significantly improved in applications using large (>50GB) heaps containing many large objects.
* Go 1.10, Many applications should experience significantly lower allocation latency and overall performance overhead when the garbage collector is active.

## GcTrace

```go
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
```

## Stack vs Heap

_"The stack is for data that needs to persist only for the lifetime of the
function that constructs it, and is reclaimed without any cost when the function
exits. The heap is for data that needs to persist after the function that
constructs it exits, and is reclaimed by a sometimes costly garbage collection."
- Ayan George

## Links

### Stacks

[Contiguous Stack Proposal](https://docs.google.com/document/d/1wAaf1rYoM4S4gtnPh0zOlGzWtrZFQ5suE8qr2sD8uWQ/pub)

### Escape Analysis and Inlining

[Go Escape Analysis Flaws](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw)
[Compiler Optimizations](https://github.com/golang/go/wiki/CompilerOptimizations)
