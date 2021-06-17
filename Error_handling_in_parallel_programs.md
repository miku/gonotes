# Error handling in parallel programs

* Create a new type that encapsulates result and error
* use package "x/sync/errgroup", [x/ErrorErrgroup](x/ErrorErrgroup)
* example: aggregate errors:
  https://github.com/kubernetes/apimachinery/blob/06deae5c9c2c030d771a467e086b6c791e8800dc/pkg/util/errors/errors.go#L231-L246,
  [x/ErrorAggregate](x/ErrorAggregate)

