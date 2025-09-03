package itertools

import (
	"iter"
)

func Repeat[V any](v V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			if !yield(v) {
				return
			}
		}
	}
}
