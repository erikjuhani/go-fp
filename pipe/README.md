# Pipe

Pipe is a mechanism or an operator that allows you to compose functions
together in a sequence, passing the output of one function as the input to the
next function. Function composition enables you to build complex computations
by chaining together simpler functions, making your code more concise and
readable.

The pipe operator is often denoted as `|>` or a similar symbol. In `go-fp` the
pipe is represented as a function `Pipe`. Pipe operator takes the result of the
expression from left side and passes it as the first argument to the function
on its right side. This process continues in a sequence as long as there are
functions to be applied.

## Usage

To create a composition call PipeN, the N as the amount of given function
arguments. The initial value is passed when calling the composed function. The
type for the initial value is inferred from the first function argument.

```go
composed := Pipe3(fn1, fn2, f3)
```

## Example

```go
func addOne(x int) int {
	return x + 1
}

func double(x int) int {
	return x * 2
}

func square(x int) int {
	return x * x
}

// Applying functions in a sequence without using pipe 
square(double(addOne(1))) // 16

// The function composition becomes harder and harder to read as we add more
// functions to the sequence
addOne(double(addOne(square(double(addOne(1)))))) // 35

// With Pipe the function composition is more human readable.
// The initial value is passed when calling the composed function
pipe.Pipe3(addOne, double, square)(1) // 16

// With Pipe the long composition does not loose readability
pipe.Pipe6(
  addOne,
  double,
  square,
  addOne,
  double,
  addOne,
)(1) // 35
```
