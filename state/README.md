# State monad

The State monad is a programming construct often used in functional programming
languages to manage stateful computations in a pure and immutable manner. It
provides a way to encapsulate stateful operations while maintaining referential
transparency. State monad represents computations that take an initial state as an input and
returns the result with the updated state as a tuple `(Result, State)`. 

State monad provides a way to work with stateful computations in a purely
functional manner by encapsulating the state and allowing for composition and
transformation of these computations.

Ultimately State monad enables to write code that manages state effectively,
while maintaining immutability and referential transparency.

## Usage

To create stateful computations we need to define the initial state `a` first
to run computations on. Initial state `s` can be anything, but it must be
wrapped into a State monad.

After wrapping the state into the State monad we can run stateful computations
on it for exampel by using `state.Modify` and `state.Put`.

Before running stateful operations we need to Run the state to access the
underlying state operation.

```go
pipe.Pipe2(state.Run[int], state.Map[int](increment))(initialState)
```

## Example

```go
// Let's create a counter using State monad.
// Here we use four state functions `Run`, `Fmap`, `Put` and `Exec`.
// Fmap is our bind function to bind the monad operations.
// Run is used to gain access to the state monad state.
// Put is used to change the actual state by setting a new state.
// Exec is used to only return the final state and not the computation.
// Essentially this function composition creates a signature func(int) int.
counter := pipe.Pipe3(
    state.Run[int], // Accessor for the state monad, starts the operation
    state.Fmap(func(s int) State[Void, int] { return Put(s + 1) }) // Increment the number by one
    state.Exec[Void](0), // Run Exec to return only the transformed final state. Pass in 0 that represents "default or zero value of the type"
)

counter(1) // Increments by one so the returned result is 2.
```
