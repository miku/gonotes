package main

import "testing"

func TestNotReally(t *testing.T) {
	a()
}

func TestA(t *testing.T) {
	if exp, result := 42, a(); exp != result {
		t.Fatalf("got %v, want %v", result, exp)
	}
}

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a()
	}
}

// go test -v -bench=.
// === RUN   TestNotReally
// --- PASS: TestNotReally (0.01s)
// === RUN   TestA
// --- PASS: TestA (0.01s)
// goos: linux
// goarch: amd64
// pkg: github.com/miku/gogwks/x/TestingBenchmark
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkA
// BenchmarkA-8         100          10397034 ns/op
// PASS
// ok      github.com/miku/gogwks/x/TestingBenchmark       1.075s
//

// $ go test -run none -bench . -benchtime 3s -benchmem
// goos: linux
// goarch: amd64
// pkg: github.com/miku/gogwks/x/TestingBenchmark
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkA-8         345          10558820 ns/op               0 B/op          0 allocs/op
// PASS
// ok      github.com/miku/gogwks/x/TestingBenchmark       4.696s
//
