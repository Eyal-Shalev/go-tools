package itertools

import (
	. "github.com/Eyal-Shalev/itertools-go/internal"
	"iter"
)

func Reject[V any](seq iter.Seq[V], f func(V) bool) iter.Seq[V] {
	return Select(seq, Not1(f))
}

func Reject2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return Select2(seq, Not2(f))
}
