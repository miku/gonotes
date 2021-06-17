---
id: Solid
ref:
- https://dave.cheney.net/2016/08/20/solid-go-design
---

# SOLID Principles

## Single Responsibility

* more on the package level

## Open-Closed

* open for extension
* closed for modification

A subtype can call supertype methods, but the supertype does not have access to
the subtype.

## Liskov Subsitition

* subclasses (types embedding another type) cannot be substituted for the superclass
* if you need substitution, use an interface
* interfaces use structural typing (as opposed to nominal typing)

Example of a large set of substitutes (implementations):

> Because io.Reader‘s deal with anything that can be expressed as a stream of
> bytes, we can construct readers over just about anything; a constant string,
> a byte array, standard in, a network stream, a gzip’d tar file, the standard
> out of a command being executed remotely via ssh.


> Require no more, promise no less.

## Interface Segregation

> Clients should not be forced to depend on methods they do not use.

* limit what arguments are passed in, e.g. `*os.File` vs `io.Reader`
* reduces the surface, allows for easier modification

> A great rule of thumb for Go is accept interfaces, return structs.

## Dependency Inversion

> High-level modules should not depend on low-level modules. Both should depend
> on abstractions. Abstractions should not depend on details. Details should
> depend on abstractions.


