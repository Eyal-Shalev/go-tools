package slicetools

import "github.com/Eyal-Shalev/go-tools/functools"

func Map[S ~[]From, From, To any](s S, f func(From) To) []To {
	out := make([]To, len(s))
	for i, val := range s {
		out[i] = f(val)
	}
	return out
}

func ToAny[S ~[]T, T any](s S) []any {
	return Map(s, functools.ToAny)
}
