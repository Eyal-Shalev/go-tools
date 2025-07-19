package itertools

import "iter"

func Map[V, T any](seq iter.Seq[V], applyFn func(V) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if !yield(applyFn(v)) {
				return
			}
		}
	}
}

func Map2[K, V, T any](seq iter.Seq2[K, V], applyFn func(V) T) iter.Seq2[K, T] {
	return func(yield func(K, T) bool) {
		for k, v := range seq {
			if !yield(k, applyFn(v)) {
				return
			}
		}
	}
}
