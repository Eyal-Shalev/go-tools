package slicetools

import "github.com/Eyal-Shalev/go-tools/functools"

func Reject[V any, S ~[]V](s S, f func(V) bool) S {
	return Select(s, functools.Not1(f))
}
