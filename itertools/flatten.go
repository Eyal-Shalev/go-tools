package itertools

import "iter"

func Flatten[T any](s iter.Seq[iter.Seq[T]]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for innerS := range s {
			for v := range innerS {
				if !yield(v) {
					return
				}
			}
		}
	}
}
