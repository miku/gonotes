# Services

## Handler

There are a couple of concepts related to services.

* [http.Handler](https://golang.org/pkg/net/http/#Handler)

The core interface for handline HTTP request.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

## Handlerfunc

You can turn a function (with the right signature) into a type, that satisfies
http.Handler.

* [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc)

> The HandlerFunc type is an adapter to allow the use of ordinary functions as
> HTTP handlers. If f is a function with the appropriate signature,
> HandlerFunc(f) is a Handler that calls f.

With first class functions, it is possible to start an HTTP handler with a
single (long) line.

## HandleFunc

If we type `http.HandleFunc`, we are working on the `DefaultServeMux`, a
[http.ServeMux](https://golang.org/pkg/net/http/#ServeMux) added for convenience
in the http package.

> ServeMux is an HTTP request multiplexer. It matches the URL of each incoming
> request against a list of registered patterns and calls the handler for the
> pattern that most closely matches the URL.

The `ServeMux` is itself an HTTP handler, as it implements
[ServeHTTP](https://golang.org/pkg/net/http/#ServeMux.ServeHTTP).


## Code Review

* [Example 1](x/Services-00-server/main.go)
* [Example 2](x/Services-01-server-handlerfunc/main.go)
* [Example 3](x/Services-02-server-handler/main.go)
* [Example 4](x/Services-03-server-handle)
* [Example 4](x/Services-04-server-handlefunc/main.go)
* [Example 5](x/Services-05-two-handlers/main.go)
* [Example 6](x/Services-06-handler-methods/main.go)
* [Example 7](x/Services-07-methods/main.go)
* [Example 8](x/Services-08-post/main.go)
* [Example 9](x/Services-09-pattern/main.go)
* [Example 10](x/Services-10-mux/main.go)
* [Example 11](x/Services-11-gorilla-mux/main.go)

## Exercises

* [Debug Handler](x/Services-12-template/main.go) ([Answer](x/Services-12-exercise/main.go))
