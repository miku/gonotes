# Methods

* can be defined on any type, not just structs (except a pointer or an interface)
* on structs, they are kind of syntactic sugar (instead of using a first argument to a function)
* functions are first class objects

You can combine first class methods and e.g. structs to carry state.

Example:

* [x/MethodCarryState](x/MethodCarryState)

## Nil is a valid receiver

Example:

* [x/MethodNil](x/MethodNil)
