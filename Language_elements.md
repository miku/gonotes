# Language elements

* Go has 25 keywords (including `goto`): [Keywords](https://golang.org/ref/spec#Keywords)
* Personal favorites: `switch`

## Mutable types

* arrays and slices
* maps
* channels
* structs

## Immutable types


* interfaces
* booleans, numeric values (including values of type int)
* strings
* pointers
* function pointers, and closures which can be reduced to function pointers
* structs having a single field

## Size of Types

Use:

```
var s = "hello"
unsafe.Sizeof(s) 
```

* https://play.golang.org/p/4mzdOKW6uQ

### Strings

> A string is represented in memory as a 2-word structure containing a pointer
> to the **string data and a length**. Because the string is immutable, it is safe
> for multiple strings to share the same storage, so slicing s results in a new
> 2-word structure with a potentially different pointer and length that still
> refers to the same byte sequence.


> This means that slicing can be done without allocation or copying, making
> string slices as efficient as passing around explicit indexes.

A gotcha:

Original string storage after slicing stays in memory; even if not needed anymore.


## Passing Values

You can pass values or pointer. Values get copied, pointers allow to mutate.

## Stack vs Heap

The go runtime can make decisions on where to store a value.

You can inspect that using escape analysis.

```
$ go build -gcflags '-m -l'
```

The tool subcommand allows you - among other things - to control the compile process, e.g. you can
inspect the assembly.

```
$ go tool compile "-m" main.go
```

* "print optimization decisions"


## Selected notes on ref/spec

A few picks from [ref/spec](https://golang.org/ref/spec).

### Semicolons

Even if you do not see them typically, Go uses
  [semicolons](https://golang.org/ref/spec#Semicolons) - they just get
  automatically inserted into the token stream.

### Identifiers

You can name a variable `αβ` - *first character in an identifier must be a letter*.

> Go treats all characters in any of the Letter categories Lu, Ll, Lt, Lm, or Lo as Unicode letters [...]

* [Unicode 4.5 General Category](https://www.unicode.org/versions/Unicode8.0.0/ch04.pdf#page=17)


### Rune literals

A rune is an alias for int32, representing a Unicode code point.

It is written with single quotes.

> A rune literal is expressed as one or more characters enclosed in single quotes, as in 'x' or '\n'.

[embedmd]:# (../x/runevalue/main.go)
```go
package main

import "fmt"

func main() {
	a := 'a'
	fmt.Printf("%T %c %d %x %v\n", a, a, a, a, a == 97)
	fmt.Printf("%s %d\n", "a", len("a"))

	v := '⧉'
	fmt.Printf("%T %c %d %x %v\n", v, v, v, v, v == 10697)

	// int32 a 97 61 true
	// int32 ⧉ 10697 29c9 true
}
```

More on that in [https://blog.golang.org/strings](https://blog.golang.org/strings).
