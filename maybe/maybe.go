// Maybe Monad represents the presence `Just` or an absence `Nothing` of a
// value. It is commonly used to handle computations that may or may not return
// a result, or computations that can potentially fail.
//
// Maybe monad is particularly useful when dealing with operations that can return
// `nil` values, as it provides a way to handle such cases in a more structured
// and safe manner. Maybe monad removes the need for imperative and explicit
// `nil` pointer checks, which effectively make the code more noisy hiding the
// important domain logic. Maybe monad essentially enables you to focus solely
// on the "happy path".
package maybe

import (
	"reflect"
)

// Maybe monad data type representation. May or may not contain a pointer
// value. Nothing is represented as a `nil` value internally.
type Maybe[A any] struct{ val *A }

// Just is the return operation for Maybe monad that returns the representation
// of existence of a value.
func Just[A any](v A) Maybe[A] {
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		rv := reflect.ValueOf(v)
		if rv.IsNil() {
			panic("trying to insert `nil` an absent value as present `Just` value")
		}
		x := rv.Elem().Interface().(A)
		return Maybe[A]{&x}
	}
	return Maybe[A]{&v}
}

// Nothing is the return operation for Maybe monad that returns the representation
// of absence of a value.
func Nothing[A any]() Maybe[A] {
	return Maybe[A]{}
}

// From is the return operation for Maybe monad that returns either Just a or
// Nothing. Intended to be used with Go functions that return tuple as `val, ok`.
func From[A any](val A, ok ...bool) Maybe[A] {
	if is_nil(val) || (len(ok) > 0 && !ok[0]) {
		return Nothing[A]()
	}

	return Just(val)
}

// Map or the "bind" function takes the contents of the Maybe monad and passes
// it to function `f` as a parameter. The function `f` returns a new Maybe
// monad as a result.
func Map[A, B any](f func(A) B) func(Maybe[A]) Maybe[B] {
	return func(m Maybe[A]) Maybe[B] {
		if m.val != nil {
			return Just(f(*m.val))
		}
		return Nothing[B]()
	}
}

// Fmap or also known as `bind` function lets non-monadic function `f` to
// operate on the contents of monad m a, and lifts the value to a new domain
// (Maybe a -> Maybe b).
func Fmap[A, B any](f func(A) Maybe[B]) func(Maybe[A]) Maybe[B] {
	return func(m Maybe[A]) Maybe[B] {
		if m.val != nil {
			return f(*m.val)
		}
		return Nothing[B]()
	}
}

// Match matches Maybe monad depending of it's current state and returns the
// value determined by the return type of b.
func Match[A, B any](Nothing func() B, Just func(A) B) func(Maybe[A]) B {
	return func(m Maybe[A]) B {
		if m.val != nil {
			return Just(*m.val)
		}
		return Nothing()
	}
}

// internal
func is_nil(v any) bool {
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		return reflect.ValueOf(v).IsNil()
	}

	return false
}
