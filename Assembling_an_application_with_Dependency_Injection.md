# Wire

Dependency injection with Google Wire.

# Intro

* two approaches: runtime and compile time

Dependency Injection introduces a new entity (e.g. a container) that helps
creating objects and their dependencies.

It comes in various flavors, like

> Constructor Injection, Setter Injection, and Interface Injection. --
> https://martinfowler.com/articles/injection.html

Some languages seem to use DI more than others:

* [https://stackoverflow.com/questions/2461702/why-is-ioc-di-not-common-in-python](https://stackoverflow.com/questions/2461702/why-is-ioc-di-not-common-in-python)

[Does DI need
interfaces?](https://stackoverflow.com/questions/43079277/do-we-need-interfaces-for-dependency-injection)
-- No, but it's more useful with interfaces.


## Implementations

Various implementations available, e.g. from
[Google](https://github.com/google/wire),
[Uber](https://github.com/uber-go/dig) and others.


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
