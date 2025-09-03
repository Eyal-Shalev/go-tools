package itertools

import (
	"iter"
)

func Tee[V any](seq iter.Seq[V], bufferSize int) (iter.Seq[V], iter.Seq[V]) {
	leftChan := make(chan V, bufferSize)
	rightChan := make(chan V, bufferSize)

	go func() {
		defer close(leftChan)
		defer close(rightChan)
		for v := range seq {
			leftChan <- v
			rightChan <- v
		}
	}()

	left := func(yield func(V) bool) {
		for v := range leftChan {
			if !yield(v) {
				return
			}
		}
	}
	right := func(yield func(V) bool) {
		for v := range rightChan {
			if !yield(v) {
				return
			}
		}
	}

	return left, right
}
