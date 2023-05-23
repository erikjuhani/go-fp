# Maybe monad

Maybe monad represents the presence `Just` or an absence `Nothing` of a value.
It is commonly used to handle computations that may or may not return a result,
or computations that can potentially fail.

Maybe monad is particularly useful when dealing with operations that can return
`nil` values, as it provides a way to handle such cases in a more structured
and safe manner. Maybe monad removes the need for imperative and explicit `nil`
pointer checks, which effectively make the code more noisy hiding the important
domain logic. Maybe monad essentially enables you to focus solely on the _happy
path_.

## Usage

To use Maybe monad one must call the return operation `Just`, `Nothing` or
`From`.

`Just` is used in the Maybe monad to wrap a value and indicate that it is
present.

`Nothing` is used in the Maybe monad to indicate an absence of a value.
`Nothing` will always be nothing when it has been set.

`From` is primarily used as a helper for functions that return a tuple `(T, bool)`.
It will create a `Just` value or `Nothing` depending on the returned value `T`
or the second value `bool`, if value `T` is `nil` the Maybe monad will be
`Nothing` as the value is absent. If the second value in the tuple is `false`,
it will also be interpreted as `Nothing`.

## Example

```go
func head[T any](slice []T) maybe.Maybe[T] {
    if len(slice) > 0 {
        return maybe.Just(slice[0])
    }
    return maybe.Nothing[T]()
}

func inverse(x int) maybe.Maybe[float32] {
    if x > 0 {
        return maybe.Just(1 / float32(x))
    }

    return maybe.Nothing[float32]()
}

// Trying to get the head of an empty string array results as `Nothing`
// and continues to be nothing until the end of the function chain
pipe.Pipe2(
    head[string],
    maybe.Map(strings.ToUpper),
)([]string{}) // -> Nothing

// However when trying to access head on an array that contains items, the
// first item is returned upper cased as a Just value
pipe.Pipe2(
    head[string],
    maybe.Map(strings.ToUpper),
)([]string{"hello", "world"}) // -> Just HELLO

// Match can be used to "pattern match" Nothing or Just values. Match returns
// an actual value and can be used to extract a Just value
pipe.Pipe3(
    head[string],
    maybe.Map(strings.ToUpper),
    maybe.Match(
        func() string { return "Got nothing" },
        func(v string) string { return v },
    )
)([]string{}) // -> "Got nothing"

// Fmap can be used to flatten the Maybe monad like in this example returned by
// the inverse function. If we would use `Map` instead we would end up with
// Maybe[Maybe[Float32]] type. `inverse` function returns `Nothing` if given
// input value of zero
pipe.Pipe2(
    head[int],
    maybe.Fmap(inverse),
)([]int{0}) // -> Maybe[Float32] -> Nothing
```
