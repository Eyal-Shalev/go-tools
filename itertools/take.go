package itertools

import (
	"iter"

	"golang.org/x/exp/constraints"
)

func TakeWhile[T any](seq iter.Seq[T], predicate func(T) bool, includeLast bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			conformsToPredicate := predicate(v)
			if !conformsToPredicate && !includeLast {
				return
			}
			if !yield(v) {
				return
			}
			if !conformsToPredicate {
				return
			}
		}
	}
}

func TakeN[T any, N constraints.Integer](seq iter.Seq[T], n N) iter.Seq[T] {
	return func(yield func(T) bool) {
		if n <= N(0) {
			return
		}
		idx := N(0)
		for v := range seq {
			if !yield(v) {
				return
			}
			if idx++; idx >= n {
				return
			}
		}
	}
}
