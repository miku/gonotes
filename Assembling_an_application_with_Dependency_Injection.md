# Wire

Dependency injection with Google Wire.

# Intro

* two approaches: runtime and compile time

# Examples

* https://github.com/miku/demo/tree/main/di

# Build tags

Wire uses conditional compilation with build tags. The file which contains the
wiring should only be considered, when we use `wireinject` - which is what the
[wire](https://git.io/JZFS4) tool [sets](https://git.io/JZFSH) when generating
code (wire uses
[x/tools/go/packages](https://pkg.go.dev/golang.org/x/tools/go/packages) to
discover which files to process).

```go
//+build wireinject
```

