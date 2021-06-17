# Environment Variables

## GOPATH

GOPATH set to $HOME/go by default. It was a workspace.

Previously, it was much more important to keep code organized under GOPATH. 

> Some people found GOPATH annoying. It certainly was peculiar. Tools have been
written to ease the *Getting Started* experience.

> If you are developing in Go, long before the release of Go 1.11, you must have
> a terrible experience as I had to deal with $GOPATH.
> (https://medium.com/mindorks/create-projects-independent-of-gopath-using-go-modules-802260cdfb51)

It contains executables under bin, pkg cache for modules and source, fetched
with go get and compiled into executables.

```go
$ ls -1 ~/go
bin
pkg
src
```

> The GOPATH is one of the trickier aspects of setting up Go, but once it is set
> up, we can usually forget about it.

## GOROOT

* your local go installation

## CGO_ENABLED

* can be set to 0 to build a pure Go binary; otherwise system libraries may be
  linked (e.g. for dns lookup)

## List all environment variables

```
$ go env
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/tir/.cache/go-build"
GOENV="/home/tir/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/home/tir/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/tir/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.16.4"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/home/tir/code/miku/gogwks/go.mod"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-bu... -gno-reco..."
```