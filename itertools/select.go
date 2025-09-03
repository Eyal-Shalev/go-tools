package itertools

import "iter"

func Select[V any](seq iter.Seq[V], f func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !f(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

func Select2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !f(k, v) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
