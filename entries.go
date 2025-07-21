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

func FromEntries[K, V any](seq iter.Seq[Entry[K, V]]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for entry := range seq {
			if !yield(entry.Key, entry.Val) {
				return
			}
		}
	}
}
