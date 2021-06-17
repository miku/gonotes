
## Assignment to entry in nil map
You have to initialize a map using the make function (or a map literal) before
you can add any elements.

## Invalid memory address or nil pointer dereference
An uninitialized pointer is nil, and you can’t follow the nil pointer. If x is
nil, an attempt to evaluate *x will cause a run-time panic.

## Multiple-value in single-value context
When a function returns multiple values, you must take care of each one. You can
use the blank identifier (underscore) to ignore unwanted return values.

## Array won’t change

Arrays in Go are values: when you pass an array to a function it gets a copy of
the original array data. If you want a function to update its elements, use a
slice that refers the array.

## Shadowed variables
An identifier declared in a block may be redeclared in an inner block. The new
variable shadows the original throughout the scope of the inner block.

## Unexpected newline
In a multi-line slice, array or map literal, every line must end with a comma.
As a result, you can add and remove lines without modifying the surrounding
code.

## Immutable strings
Go strings are read-only byte slices (with a few extra properties). To update
string data, consider using a rune slice instead.

## How does characters add up?
Characters (rune literals) are integer values used to identify Unicode code
points. You can turn a code point into a string with a conversion.

## Where is my copy?
Copy copies the minimum number of elements in the destination and source slices.
To make a full copy, you must allocate a big enough destination slice.

## Why doesn’t append work every time? [scary bug]
If there is enough capacity, append reuses the underlying array.

## Constant overflows int
An untyped integer constant is converted to an int when the type can’t be
inferred from the context.

## Unexpected ++, expecting expression
In Go increment and decrement operations can’t be used as expressions, only as
statements. Only postfix notation is available.

## Get your priorities right

In Go the multiplication, division, and remainder operators have the same
precedence.

## Go and Pythagoras

The circumflex (^) denotes bitwise XOR in Go.

 ## No end in sight
    An integer overflow occurs when an arithmetic operation tries to create a value that is outside the range that can be represented.
    
## Numbers that start with zero
    Octal literals start with 0, hexadecimal with 0x, and decimal literals never start with zero.

## Whatever remains
    The remainder operator can give negative answers if the dividend is negative.

## Time is not a number
    There is no mixing of numeric types in Go.
    
## Index out of range
    Arrays, slices and strings are indexed starting from zero.

## Unexpected values in range loop
    The range loop generates two values: first the index, then the data.

## Can’t change entries in range loop
    The range loop uses a local variable to store iteration values.
    
## Iteration variable doesn’t see change in range loop
    The range expression is evaluated once before beginning the loop.

## Iteration variables and closures
    A data race occurs when two goroutines access the same variable concurrently and at least one of the accesses is a write.

## No JSON in sight
    Only the the exported fields of a Go struct will be present in the JSON output.

## Nil is not nil
    An interface value is nil only if the concrete value and dynamic type are both nil.
