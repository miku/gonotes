# Error Handling

* go has no exceptions, but multiple return values
* error is an predeclared (interface) type

```
type error interface {
	Error() string
}
```

The absence of error is marked by a nil value.

> An error value may be of any type which satisfies the language-defined error
> interface. A program can use a type assertion or type switch to view an error
> value as a more specific type.

Example from the standard library:

```go
// PathError records an error and the operation and file path that caused it.
type PathError struct {
        Op   string
        Path string
        Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

func (e *PathError) Unwrap() error { return e.Err }

// Timeout reports whether this error represents a timeout.
func (e *PathError) Timeout() bool {
        t, ok := e.Err.(interface{ Timeout() bool })
        return ok && t.Timeout()
}
```

A few problems.

> Creating a new error with fmt.Errorf discards everything from the original
> error except the text. 

This is what the newer `Unwrap() error` method is for; if provided by the error
type, it allows to get an wrapper error value.

For examination, two methods were added in Go 1.13:

* `errors.Is`, e.g. `errors.Is(err, ErrPermission)`
* `errors.As`

```go
// Similar to:
//   if e, ok := err.(*QueryError); ok { … }
var e *QueryError
// Note: *QueryError is the type of the error.
if errors.As(err, &e) {
    // err is a *QueryError, and e is set to the error's value
}
```

Finally, a new verb is available on `fmt.Errorf` - `%w` - as a shortcut to
support unwrap.

```go
if err != nil {
    // Return an error which unwraps to err.
    return fmt.Errorf("decompress %v: %w", name, err)
}
``` 

Further customization of a new error type is possible, e.g. one can implement
`Is(target error) bool` in order to control the comparison.

## Panic

Is an exception, but not the main way to handle errors. It comes with a
compangion function `recover`.

> Recover is a built-in function that regains control of a panicking goroutine.
> Recover is only useful inside deferred functions. During normal execution, a
> call to recover will return nil and have no other effect. 