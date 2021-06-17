# Slices

* three values, pointer, length, capacity
* Slice tricks: https://github.com/golang/go/wiki/SliceTricks

## Relevant builtin functions

The builtin `append` can be used to append one element or a slice to an existing
slice. Builtin `copy` to duplicate a slice.

In summary:

* make, new -- [x/SliceMakeVsNew](x/SliceMakeVsNew)
* len
* cap
* append
* copy

> The zero value of a slice is nil. The len and cap functions will both return 0 for a nil slice.

Refs:

* [builtin](https://golang.org/pkg/builtin/)

## From Slice to Array and back

* [x/SliceStorage](x/SliceStorage)

## Examples

What is happening?

* [A] [SliceCut](x/SliceCut) 

More examples:

* [SliceBatching](x/SliceBatching)
* [SliceCut](x/SliceCut)
* [SliceCutSafe](x/SliceCutSafe)
* [SliceDeleteSafe](x/SliceDeleteSafe)
* [SliceExpand](x/SliceExpand)
* [SliceExtend](x/SliceExtend)
* [SliceFilter](x/SliceFilter)
* [SliceFilterNoAlloc](x/SliceFilterNoAlloc)

## [A] Slice Gotchas

What is happening in the following snippet.

* [x/GotchaAppend](x/GotchaAppend)
* [x/GotchaAppendBytes](x/GotchaAppendBytes)

## Examples of slice operations

* [x/SliceExtend](x/SliceExtend)
* [x/SliceDeleteSafe](x/SliceDeleteSafe)
* [x/SliceCutSafe](x/SliceCutSafe)

