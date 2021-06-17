# HTTP requests

The [net/http](https://golang.org/pkg/net/http/) package contains HTTP server and client code.

## Making a GET request

The package comes with a default client, that is available directly - it will
return a `http.Response` and an error.

```
resp, err := http.Get("http://example.com/")
```

## Code Review

* [Example 1](x/Http1)
* [Example 2](x/Http2)
* [Example 3](x/Http3)
* [Example 4](x/Http4)
* [Example 5](x/Http5)

## Exercises

