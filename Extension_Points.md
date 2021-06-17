# Extension Points

An extension point provides an opportunity to customize or adapt an existing
facility to new requirements.

# Custom Flags

Extending flag processing, e.g. to support repeated flags.

[embedmd]:# (x/flagrepeat/flag.go)
```go
// Package xflag add an additional flag type Array for repeated string flags.
//
//   var f xflag.Array
//   flag.Var(&f, "r", "some repeatable flag")
//
//   flag.Parse()                // $ command -r a -r b -r c
//   for _, v := range f { ... } // []string{"a", "b", "c"}
//
package xflag

import "strings"

// ArrayFlags allows to store lists of flag values.
type Array []string

// String representation.
func (f *Array) String() string {
	return strings.Join(*f, ", ")
}

// Set appends a value.
func (f *Array) Set(value string) error {
	*f = append(*f, value)
	return nil
}
```

# Implementing a custom io.Reader

* Go Tour [methods/21](https://tour.golang.org/methods/21),
  [methods/22](https://tour.golang.org/methods/22): Readers

Example for a small interface with many different implementations.


# Custom Struct Tags

Struct tags allow you to annotate a struct. One application is serialization.
With a little helper library, it is relatively easy to implement your own
struct tags.

Example: TSV reading

* [https://github.com/miku/demo/encoding](https://github.com/miku/demo/encoding)

---
id: CustomRoundtripper
refs:
- https://medium.com/@_jesus_rafael/making-http-client-more-resilient-in-go-d24c66a64bd1
---

# Custom Roundtripper

Smaller interfaces can help writing alternative implementations quickly. The
[http.Client](https://golang.org/pkg/net/http/#Client) delegates the HTTP
transaction to a
[http.Roundtripper](https://golang.org/pkg/net/http/#RoundTripper), a single
method interface, which

> executes a single HTTP transaction, returning a Response for the provided Request.

A few alternative roundtrippers:

* [https://github.com/gregjones/httpcache](https://github.com/gregjones/httpcache);
  allows to cache HTTP in a mostly
[RFC7234](https://datatracker.ietf.org/doc/html/rfc7234) compliant way.  HTTP
caching can be implemented in a variety of ways, e.g. also outside the http
package. However, this approach uses an extension point in the library (the
http.RoundTripper interface) and implements a solution close to the problem.
Client code gets simpler, too: one only needs to change the client transport.




