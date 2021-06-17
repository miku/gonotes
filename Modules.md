# Modules

A newer mechanism to deal with dependencies.

## Basics

Run `go mod init` to start a project. This will add a `go.mod` file to your
directory, which should be version controlled.

The `go get` command will update this file (as well as place executables on the
system).

> After go get updates the go.mod file, it builds the packages named on the
> command line. Executables will be installed in the directory named by the
> GOBIN environment variable, which defaults to $GOPATH/bin or $HOME/go/bin if
> the GOPATH environment variable is not set.

New in Go 1.16:

> Since Go 1.16, go install is the recommended command for building and
> installing programs. When used with a version suffix (like @latest or
> @v1.4.6), go install builds packages in module-aware mode, ignoring the go.mod
> file in the current directory or any parent directory, if there is one.

## Edit workflow

* edit code, add import

Previously, go.mod would have been automatically updated by go get - but this
has been "fixed".

``` 
$ go run main.go 
main.go:3:8: no required module provides package github.com/fatih/color; to add it:
        go get github.com/fatih/color
```

After go get go.mod has been updated:

```
module github.com/miku/demo

go 1.16

require (
        github.com/fatih/color v1.12.0 // indirect
        github.com/google/wire v0.5.0 // indirect
)
```

Sync with source code with `go mod tidy`:

```
$ go mod tidy
$ cat go.mod

module github.com/miku/demo

go 1.16

require (
        github.com/fatih/color v1.12.0
        github.com/google/wire v0.5.0
)
```

Change revision. If the revision does not exist:

```
$ go mod tidy
go: github.com/fatih/color@v1.100.0: reading github.com/fatih/color/go.mod at revision v1.100.0: unknown revision v1.100.0
```

## What is a version

* a git tag
* a combination of commit hash and date, if not tag is defined

## Fixing version

* edit go.mod file and set another tag

## Freeze your dependencies

You can run `go mod vendor` to vendor code under the `vendor` directory, e.g. to
ensure reproducibilty.

> When vendoring is enabled, the go command will load packages from the vendor
> directory instead of downloading modules from their sources into the module
> cache and using packages those downloaded copies.

> When using modules, the go command typically satisfies dependencies by
> downloading modules from their sources into the module cache, then loading
> packages from those downloaded copies. Vendoring may be used to allow
> interoperation with older versions of Go, or to ensure that all files used for
> a build are stored in a single file tree.

> If the vendor directory is present in the main module's root directory, it
> will be used automatically if the go version in the main module's go.mod file
> is 1.14 or higher. To explicitly enable vendoring, invoke the go command with
> the flag -mod=vendor. To disable vendoring, use the flag -mod=readonly or
> -mod=mod.

## List outdate modules

```
$ go list -m -u all
github.com/syndtr/goleveldb
github.com/fsnotify/fsnotify v1.4.7 [v1.4.9]
github.com/golang/protobuf v1.2.0 [v1.5.2]
github.com/golang/snappy v0.0.1 [v0.0.3]
github.com/hpcloud/tail v1.0.0
github.com/onsi/ginkgo v1.7.0 [v1.16.4]
github.com/onsi/gomega v1.4.3 [v1.13.0]
golang.org/x/net v0.0.0-20180906233101-161cd47e91fd [v0.0.0-20210614182718-04defd469f4e]
golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f [v0.0.0-20210220032951-036812b2e83c]
golang.org/x/sys v0.0.0-20180909124046-d0be0721c37e [v0.0.0-20210616094352-59db8d763f22]
golang.org/x/text v0.3.0 [v0.3.6]
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 [v1.0.0-20201130134442-10cb98267c6c]
gopkg.in/fsnotify.v1 v1.4.7
gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
gopkg.in/yaml.v2 v2.2.1 [v2.4.0]
```