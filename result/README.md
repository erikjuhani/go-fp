# Result monad

Result monad encapsulates two possible values: a successful result `Ok` or an
error `Err`. Result monad allows to chain together a series of operations that
can potentially fail and propagate the error through the computations. This
enables to write code that is more focused on the happy path and separate the
error handling logic in a clean and concise way.

When a computation succeeds, Result monad carries the successful value and
allows you to continue chaining operations on it. If at any point an error
occurs, the error value is propagated, and subsequent operations are bypassed
until an error handler is encountered, which enables explicit handling of both
the success and error cases making the code easier to reason about.

Result monad is similar to Maybe monad, the main difference is that it holds
either the value or the error, and not value or nothing.

## Usage

To use Result monad one must call the return operation `Ok`, `Err` or
`From`.

`Ok` is used in the Result monad to wrap a value and indicate that it is
successful.

`Err` is used in the Result monad to indicate a failure state with an error
message. Result monad will always be in failure state when a failure state has
been set.

`From` is primarily used as a helper for functions that return a tuple `(T, error)`.
It will create an `Ok` value or failure `Err` depending on the returned error
value `error`.

## Example

```go
func head[T any](slice []T) result.Result[T] {
    if len(slice) > 0 {
        return result.Ok(slice[0])
    }
    return result.Err[T](errors.New("cannot get head from an empty array"))
}

func inverse(x int) result.Result[float32] {
    if x == 0 {
        return result.Err[float32](errors.New("division by zero"))
    }

    return result.Ok(1 / float32(x))
}

// Trying to get the head of an empty string array results as a failure `Err`
// and continues to be in failure state until the end of the function chain
pipe.Pipe2(
    head[string],
    result.Map(strings.ToUpper),
)([]string{}) // -> Err "cannot get head from an empty array"

// However when trying to access head on an array that contains items, the
// first item is returned upper cased as an `Ok` value
pipe.Pipe2(
    head[string],
    result.Map(strings.ToUpper),
)([]string{"hello", "world"}) // -> Ok HELLO

// Match can be used to "pattern match" Err or Ok values. Match returns
// an actual value and can be used to extract an Ok value
pipe.Pipe3(
    head[string],
    result.Map(strings.ToUpper),
    result.Match(
        func(err error) string { return err.Error() },
        func(val string) string { return val },
    )
)([]string{}) // -> "cannot get head from an empty array"

// Fmap can be used to flatten the Result monad like in this example returned
// by the `inverse` function. If we would use `Map` instead we would end up
// with Result[Result[Float32]] type. `inverse` function returns `Err` if given
// input value of zero
pipe.Pipe2(
    head[int],
    result.Fmap(inverse),
)([]int{0}) // -> Result[Float32] -> Err "division by zero"
```
