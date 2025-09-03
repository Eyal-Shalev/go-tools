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

func Map2[K, V, KOut, VOut any](seq iter.Seq2[K, V], applyFn func(K, V) (KOut, VOut)) iter.Seq2[KOut, VOut] {
	return func(yield func(KOut, VOut) bool) {
		for k, v := range seq {
			if !yield(applyFn(k, v)) {
				return
			}
		}
	}
}
