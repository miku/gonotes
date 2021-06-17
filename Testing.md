# Testing

## Conventions

* test files live beside code
* `_test.go`

## Running a test

Various selectors:

```
$ go test .
```

Run all test in current package.

```
$ go test ./...
```

Run tests verbose:

```
$ go test -v ./...
```

## Defining a Test and Benchmark

Example:

* [x/TestingBenchmark](x/TestingBenchmark)


## Compiling a test

```
$ go test -c
```

You can stress test a package (via:
[https://dave.cheney.net/2013/06/19/stress-test-your-go-packages](https://dave.cheney.net/2013/06/19/stress-test-your-go-packages))

```bash
#!/bin/bash -e

go test -c -race
PKG=$(basename $(pwd))

while true ; do
    export GOMAXPROCS=$[ 1 + $[ RANDOM % 128 ]]
    ./$PKG.test $@ 2>&1
done

```