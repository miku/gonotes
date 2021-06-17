# Importing code

Typically, code will be imported and the last path element will serve as a
qualifier to access exported identifiers.

There are a few variants:

* normal/default
* aliasing
* period (.)
* blank (for side effects)

```
import   "lib/math"         math.Sin
import m "lib/math"         m.Sin
import . "lib/math"         Sin

import(
  "image"
  "image/jpeg"   // direct access
  _ "image/png"  // allow to read png
  _ "image/gif"  // allow to read gif
)
```

> Decoding any particular image format requires the prior registration of a
> decoder function. Registration is typically automatic as a side effect of
> initializing that format's package.


## Naming conventions

* Why is a package like `https://github.com/json-iterator/go` imported as `jsoniter`.
* What about versioning, e.g. https://github.com/googleapis/google-api-go-client

For youtube, the package declaration reads like:

* https://github.com/googleapis/google-api-go-client/blob/master/youtube/v3/youtube-gen.go

```go
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package youtube // import "google.golang.org/api/youtube/v3"
...
```

We see comment behind the package name; this is a vanity URL package and the
comment enforces the use of one name (via vanity URL).


## Vanity

```
$ sudo nc -l 80
GET /?go-get=1 HTTP/1.1
Host: 127.0.0.1
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip

----

$ GOINSECURE=127.0.0.1 GOPROXY="" go get 127.0.0.1/foo
```

For the youtube example:

```
$ curl -sL "google.golang.org/api/youtube/v3/?go-get=1" | grep "meta"
<meta name="go-import" content="google.golang.org/api git https://github.com/googleapis/google-api-go-client">
<meta name="go-source" content="google.golang.org/api https://github.com/googleapis/google-api-go-client https://github.com/googleapis/google-api-go-client/tree/master{/dir} https://github.com/googleapis/google-api-go-client/tree/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://pkg.go.dev/google.golang.org/api/youtube/v3/">
```

> In go, it is convention that the directory match the package name but this
> doesn't have to be the case, and often it is not with 3rd party packages. The
> stdlib does do a good job of sticking to this convention.


## Example

* https://github.com/miku/demo/cmd/impv