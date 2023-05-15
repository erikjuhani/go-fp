<table align="center">
  <td>

```
               __       
   __ _  ___  / _|_ __  
  / _  |/ _ \| |_|  _ \ 
 | (_| | (_) |  _| |_) |
  \__  |\___/|_| | .__/ 
  |___/          |_|    
```

  </td>
</table>

The `go-fp` library provides functional programming concepts for Go programming
language. The library itself started as a research and exploration effort in
functional programming patterns in a language that is by design imperative. The
library heavily utilizes [generics](https://tip.golang.org/doc/go1.18#generics)
introduced in Go version `1.18`.

## Background

The Go programming language was a great choice, in my opinion, to understand the underlying
functionality of function composition and patterns. The language itself is very
minimal and does not provide higher concepts of functional programming out of
the box. As the author of this library I strongly feel that in order to learn
something fundamentally one must do actual hands on work.

At first this was only meant to serve as an exploration and a study, but I
thought that others could find it useful, at least on a conceptual level.

## Installation

To use the `go-fp` library, you need to have Go installed and set up on your
system. You can install the library using the following command:

```sh
go get github.com/erikjuhani/go-fp
```

## Inspiration

Here's a list of existing libraries that provide functional concepts for Go
programming language and served as an inspiration for go-fp library.

- [fp-go](https://github.com/repeale/fp-go)
- [go-functional](https://github.com/BooleanCat/go-functional)

Other inspiration sources were F#, Haskell and fp-ts.

## References

- [haskell-wiki](https://wiki.haskell.org)

## License

[MIT](/LICENSE)
