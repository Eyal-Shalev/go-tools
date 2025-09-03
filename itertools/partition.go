package itertools

import "iter"

func Partition[V any](seq iter.Seq[V], bufferSize int, f func(V) bool) (iter.Seq[V], iter.Seq[V]) {
	a, b := Tee(seq, bufferSize)
	trueSeq := func(yield func(V) bool) {
		for v := range a {
			if !f(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
	falseSeq := func(yield func(V) bool) {
		for v := range b {
			if f(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}

	return trueSeq, falseSeq
}
