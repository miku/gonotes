# Structs

## Notes

* the empty struct consumes no storage

That's why it popular in set implementations based on a map:

```
var m map[string]struct{}
```

Example: https://play.golang.org/p/PyGYFmPmMt

Declaring many empty struct still consumes resources, e.g. compile time.

* A zero value struct is simply a struct variable where each key's value is set
  to their respective zero value.

* Conversion when tags differ? Example:
  [x/StructTagConvert](x/StructTagConvert); discussed in
  [2016](https://github.com/golang/go/issues/16085).