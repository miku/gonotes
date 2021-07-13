# Opaque Types

Typically, you will have both exported and unexported types in a package. Each
type can have exported an unexported fields.

An opaque type hides parts of its implementation, hence allowing it to change
the implementation in the future without having to care about backwards
compatibility (as far as the unexported fields are concerned).

## Example: time.Time

* Included in [Go 1.9](https://golang.org/doc/go1.9#monotonic-time), as "Transparent Monotonic Time support"
* Change proposal: [https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md](https://go.googlesource.com/proposal/+/master/design/12914-monoto# Opaque Types

Typically, you will have both exported and unexported types in a package. Each
type can have exported an unexported fields.

An opaque type hides parts of its implementation, hence allowing it to change
the implementation in the future without having to care about backwards
compatibility (as far as the unexported fields are concerned).

## Example: time.Time

* Included in [Go 1.9](https://golang.org/doc/go1.9#monotonic-time), as "Transparent Monotonic Time support"
* Change proposal: [https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md](https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md), [Discussion](https://github.com/golang/go/issues/12914)

Especially, the proposed internal change is described here:

* [https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md#time-representation](https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md#time-representation)

> Before.

```go
type Time struct {
    sec  int64     // seconds since Jan 1, year 1 00:00:00 UTC
    nsec int32     // nanoseconds, in [0, 999999999]
    loc  *Location // location, for minute, hour, month, day, year
}
```

> After (1.9 and later).

```go
type Time struct {
    wall uint64    // wall time: 1-bit flag, 33-bit sec since 1885, 30-bit nsec
    ext  int64     // extended time information
    loc  *Location // location
}
```

## Counter example: the IP type in the standard library

The underlying type of net.IP is a byte slice. It is not opaque, it's
implementation mostly cannot be changed without breaking compatibility.

A more efficient implementation lives in [inet.af/netaddr](https://github.com/inetaf/netaddr/blob/main/netaddr.go).
nic.md), [Discussion](https://github.com/golang/go/issues/12914)

Especially, the proposed internal change is described here:

* [https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md#time-representation](https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md#time-representation)

> Before.

```go
type Time struct {
    sec  int64     // seconds since Jan 1, year 1 00:00:00 UTC
    nsec int32     // nanoseconds, in [0, 999999999]
    loc  *Location // location, for minute, hour, month, day, year
}
```

> After (1.9 and later).

```go
type Time struct {
    wall uint64    // wall time: 1-bit flag, 33-bit sec since 1885, 30-bit nsec
    ext  int64     // extended time information
    loc  *Location // location
}
```

## Counter example: the IP type in the standard library

The underlying type of net.IP is a byte slice. It is not opaque, it's
implementation mostly cannot be changed without breaking compatibility.

A more efficient implementation lives in [inet.af/netaddr](https://github.com/inetaf/netaddr/blob/main/netaddr.go).
