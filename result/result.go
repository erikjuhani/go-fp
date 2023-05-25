// Result monad encapsulates two possible values: a successful result `Ok` or
// an error `Err`. Result monad allows to chain together a series of operations
// that can potentially fail and propagate the error through the computations.
// This enables to write code that is more focused on the happy path and
// separate the error handling logic in a clean and concise way.
//
// When a computation succeeds, Result monad carries the successful value and
// allows you to continue chaining operations on it. If at any point an error
// occurs, the error value is propagated, and subsequent operations are bypassed
// until an error handler is encountered, which enables explicit handling of
// both the success and error cases making the code easier to reason about.
//
// Result monad is similar to Maybe monad, the main difference is that it holds
// either the value or the error, and not value or nothing.
package result

import "reflect"

// Result monad data type representation. Contains either value `a` or an
// `error`
type Result[A any] struct {
	err error
	val A
}

// Unwrap gets the success value of the Result monad and defaults to zero value
// if Result monad contains an error
func Unwrap[A any](m Result[A]) A {
	if IsErr(m) {
		return reflect.Zero(reflect.TypeOf(m.val)).Interface().(A)
	}
	return m.val
}

// Unsafe_Unwrap unwraps the inner value of the Result monad as a tuple `(T, error)`
// It is unsafe due to panicking if Result contains error value. The panic
// message is taken from the contained error. It is highly discouraged to use
// this function
func Unsafe_Unwrap[A any](m Result[A]) A {
	if IsErr(m) {
		panic(m.err)
	}
	return m.val
}

// IsOk is a helper function for Result monad and results true if the Result
// monad contains a successful value
func IsOk[A any](m Result[A]) bool {
	return m.err == nil
}

// IsErr is a helper function for Result monad and results true if the Result
// monad contains a failure state
func IsErr[A any](m Result[A]) bool {
	return m.err != nil
}

// From is the return operation for Result monad that returns either Ok a or
// Err. Intended to be used with Go functions that return tuple as `(T, error)`
func From[A any](val A, err ...error) Result[A] {
	if len(err) == 0 || err[0] == nil {
		return Ok(val)
	}

	return Err[A](err[0])
}

// Ok is the return operation for Result monad that returns the representation
// of successful operation
func Ok[A any](val A) Result[A] {
	return Result[A]{val: val}
}

// Err is the return operation for Result monad that returns the representation
// of failing operation
func Err[A any](err error) Result[A] {
	return Result[A]{err: err}
}

// Map function takes the contents of the Result monad and passes
// it to function `f` as a parameter. The function `f` returns a new Result
// monad as the result
func Map[A, B any](f func(A) B) func(Result[A]) Result[B] {
	return func(m Result[A]) Result[B] {
		if IsErr(m) {
			return Err[B](m.err)
		}
		return Ok(f(m.val))
	}
}

// Fmap or also known as `bind` function lets non-monadic function `f` to
// operate on the contents of monad m a, and lifts the value to a new domain
// (Result a -> Result b)
func Fmap[A, B any](f func(A) Result[B]) func(Result[A]) Result[B] {
	return func(m Result[A]) Result[B] {
		if IsErr(m) {
			return Err[B](m.err)
		}
		return f(m.val)
	}
}

// Match matches Result monad depending of it's current state and returns the
// value determined by the return type of b
func Match[A, B any](Error func(err error) B, Ok func(val A) B) func(Result[A]) B {
	return func(m Result[A]) B {
		if IsErr(m) {
			return Error(m.err)
		}
		return Ok(m.val)
	}
}
