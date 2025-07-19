package itertools

import "iter"

func Enumerate[V any](seq iter.Seq[V]) iter.Seq2[uint, V] {
	idx := uint(0)
	return func(yield func(uint, V) bool) {
		for v := range seq {
			if !yield(idx, v) {
				return
			}
			idx++
		}
	}
}

func Enumerate2[K, V any](seq iter.Seq2[K, V]) iter.Seq2[uint, Entry[K, V]] {
	return Enumerate(Entries(seq))
}
