package itertools

import "iter"

type Entry[K, V any] struct {
	Key K
	Val V
}

func Entries[K, V any](seq iter.Seq2[K, V]) iter.Seq[Entry[K, V]] {
	return func(yield func(Entry[K, V]) bool) {
		for k, v := range seq {
			if !yield(Entry[K, V]{k, v}) {
				return
			}
		}
	}
}
