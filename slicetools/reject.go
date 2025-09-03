package slicetools

import "github.com/Eyal-Shalev/itertools-go/internal"

func Reject[V any, S ~[]V](s S, f func(V) bool) S {
	return Select(s, internal.Not1(f))
}
