// Provides `Pipe` function for function composition and chaining
//
// `Pipe` function composes functions in a sequence from left to right and
// takes the initial data as the first argument when invoking the resulting
// composed function, as demonstrated by `Pipe2(fn1, fn2)(initialdata)`
//
// To ensure type safety in go, it is necessary to define separate `Pipe`
// functions for each specific number of arguments. for instance, if there are
// three pipeable arguments, we would call `Pipe3(fn1, fn2, fn3)`.
package pipe

// Pipe takes one function as an argument and returns a function that takes
// a value `a` and returns the output `b` of the given function
func Pipe[A, B any](
	ab func(A) B,
) func(A) B {
	return func(a A) B { return ab(a) }
}

// Pipe2 takes two functions as arguments and returns a function that takes a
// value `a` and returns the output `c` obtained by applying the function
// arguments in a sequence
func Pipe2[A, B, C any](
	ab func(A) B,
	bc func(B) C,
) func(A) C {
	return func(a A) C { return bc(Pipe(ab)(a)) }
}

// Pipe3 takes three functions as arguments and returns a function that takes a
// value `a` and returns the output `d` obtained by applying the function
// arguments in a sequence
func Pipe3[A, B, C, D any](
	ab func(A) B,
	bc func(B) C,
	cd func(C) D,
) func(A) D {
	return func(a A) D { return cd(Pipe2(ab, bc)(a)) }
}

// Pipe4 takes four functions as arguments and returns a function that takes a
// value `a` and returns the output `e` obtained by applying the function
// arguments in a sequence
func Pipe4[A, B, C, D, E any](
	ab func(A) B,
	bc func(B) C,
	cd func(C) D,
	de func(D) E,
) func(A) E {
	return func(a A) E { return de(Pipe3(ab, bc, cd)(a)) }
}

// Pipe5 takes five functions as arguments and returns a function that takes a
// value `a` and returns the output `f` obtained by applying the function
// arguments in a sequence
func Pipe5[A, B, C, D, E, F any](
	ab func(A) B,
	bc func(B) C,
	cd func(C) D,
	de func(D) E,
	ef func(E) F,
) func(A) F {
	return func(a A) F { return ef(Pipe4(ab, bc, cd, de)(a)) }
}

// Pipe6 takes six functions as arguments and returns a function that takes a
// value `a` and returns the output `g` obtained by applying the function
// arguments in a sequence
func Pipe6[A, B, C, D, E, F, G any](
	ab func(A) B,
	bc func(B) C,
	cd func(C) D,
	de func(D) E,
	ef func(E) F,
	fg func(F) G,
) func(A) G {
	return func(a A) G { return fg(Pipe5(ab, bc, cd, de, ef)(a)) }
}

// Pipe7 takes seven functions as arguments and returns a function that takes a
// value `a` and returns the output `h` obtained by applying the function
// arguments in a sequence
func Pipe7[A, B, C, D, E, F, G, H any](
	ab func(A) B,
	bc func(B) C,
	cd func(C) D,
	de func(D) E,
	ef func(E) F,
	fg func(F) G,
	gh func(G) H,
) func(A) H {
	return func(a A) H { return gh(Pipe6(ab, bc, cd, de, ef, fg)(a)) }
}

// Pipe8 takes eight functions as arguments and returns a function that takes a
// value `a` and returns the output `i` obtained by applying the function
// arguments in a sequence
func Pipe8[A, B, C, D, E, F, G, H, I any](
	ab func(A) B,
	bc func(B) C,
	cd func(C) D,
	de func(D) E,
	ef func(E) F,
	fg func(F) G,
	gh func(G) H,
	hi func(H) I,
) func(A) I {
	return func(a A) I { return hi(Pipe7(ab, bc, cd, de, ef, fg, gh)(a)) }
}

// Pipe9 takes nine functions as arguments and returns a function that takes a
// value `a` and returns the output `j` obtained by applying the function
// arguments in a sequence
func Pipe9[A, B, C, D, E, F, G, H, I, J any](
	ab func(A) B,
	bc func(B) C,
	cd func(C) D,
	de func(D) E,
	ef func(E) F,
	fg func(F) G,
	gh func(G) H,
	hi func(H) I,
	ij func(I) J,
) func(A) J {
	return func(a A) J { return ij(Pipe8(ab, bc, cd, de, ef, fg, gh, hi)(a)) }
}

// Pipe10 takes ten functions as arguments and returns a function that takes a
// value `a` and returns the output `k` obtained by applying the function
// arguments in a sequence
func Pipe10[A, B, C, D, E, F, G, H, I, J, K any](
	ab func(A) B,
	bc func(B) C,
	cd func(C) D,
	de func(D) E,
	ef func(E) F,
	fg func(F) G,
	gh func(G) H,
	hi func(H) I,
	ij func(I) J,
	jk func(J) K,
) func(A) K {
	return func(a A) K { return jk(Pipe9(ab, bc, cd, de, ef, fg, gh, hi, ij)(a)) }
}

// Pipe11 takes ten functions as arguments and returns a function that takes a
// value `a` and returns the output `l` obtained by applying the function
// arguments in a sequence
func Pipe11[A, B, C, D, E, F, G, H, I, J, K, L any](
	ab func(A) B,
	bc func(B) C,
	cd func(C) D,
	de func(D) E,
	ef func(E) F,
	fg func(F) G,
	gh func(G) H,
	hi func(H) I,
	ij func(I) J,
	jk func(J) K,
	kl func(K) L,
) func(A) L {
	return func(a A) L { return kl(Pipe10(ab, bc, cd, de, ef, fg, gh, hi, ij, jk)(a)) }
}

// Pipe12 takes ten functions as arguments and returns a function that takes a
// value `a` and returns the output `m` obtained by applying the function
// arguments in a sequence
func Pipe12[A, B, C, D, E, F, G, H, I, J, K, L, M any](
	ab func(A) B,
	bc func(B) C,
	cd func(C) D,
	de func(D) E,
	ef func(E) F,
	fg func(F) G,
	gh func(G) H,
	hi func(H) I,
	ij func(I) J,
	jk func(J) K,
	kl func(K) L,
	lm func(L) M,
) func(A) M {
	return func(a A) M { return lm(Pipe11(ab, bc, cd, de, ef, fg, gh, hi, ij, jk, kl)(a)) }
}
