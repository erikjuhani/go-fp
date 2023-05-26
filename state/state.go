// State monad represents computations that take an initial state as an input and
// returns the result with the updated state as a tuple `(Result, State)`.
//
// State monad provides a way to work with stateful computations in a purely
// functional manner by encapsulating the state and allowing for composition and
// transformation of these computations.
//
// State monad enables to write code that manages state effectively, while
// maintaining immutability and referential transparency.
//
// The purpose of a State monad is to make stateful logic more maintainable and
// modular
package state

// State represents the state monad type, which holds a state in pure
// functional context and the state transformation itself is encapsulated
// within the monad construct
type State[A, S any] func(S) (A, S)

// Void represents an unit or void that is present in an operation that does
// not produce a result
type Void struct{}

// Run accesses the state processing function enabling to reach the function to
// operate on the state itself.
func Run[S any](s S) State[S, S] {
	return func(S) (S, S) {
		return s, s
	}
}

// Get retrieves the current state without modifying it and sets it as the
// result `(Result, State)`
func Get[S any](S) State[S, S] {
	return func(s S) (S, S) {
		return s, s
	}
}

// GetS provides a way to access and manipulate a specific state component
// without modifying the overall state itself
func GetS[A, S any](f func(S) A) State[A, S] {
	return func(s S) (A, S) { return f(s), s }
}

// Exec discards the computed result and returns only the final state. Exec is
// useful when only interested in the state transformation during a stateful
// computation.
func Exec[A, S any](s S) func(State[A, S]) S {
	return func(m State[A, S]) S {
		_, r := m(s)
		return r
	}
}

// Eval discards the final state and returns only the computed result. Eval is
// useful when only interested in the computation result.
func Eval[A, S any](s S) func(State[A, S]) A {
	return func(m State[A, S]) A {
		r, _ := m(s)
		return r
	}
}

// Map function takes the contents of the State monad and passes it to function
// `f` as a parameter. The function `f` returns a new State monad as the result
func Map[S, A, B any](f func(A) B) func(State[A, S]) State[B, S] {
	return func(m State[A, S]) State[B, S] {
		return func(s1 S) (B, S) {
			a, s2 := m(s1)
			return f(a), s2
		}
	}
}

// Fmap or also known as `bind` function lets non-monadic function `f` to
// operate on the contents of monad m a, and lifts the value to a new domain
// (State a -> State b).
func Fmap[A, B, S any](f func(A) State[B, S]) func(State[A, S]) State[B, S] {
	return func(m State[A, S]) State[B, S] {
		return func(s1 S) (B, S) {
			a, s2 := m(s1)
			return f(a)(s2)
		}
	}
}

// Modify transforms the current state based on the given function `f`. Modify
// does not yield any meaningful result, it only updates the state to a new
// state based on the given function `f`
func Modify[S any](f func(S) S) State[Void, S] {
	return func(s S) (Void, S) {
		return Void{}, f(s)
	}
}

// Put replaces the current state with a new state. Put does not yield any
// meaningful result, it only sets the state to a new provided state `a`
func Put[A any](s A) State[Void, A] {
	return func(A) (Void, A) {
		return Void{}, s
	}
}
