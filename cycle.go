package itertools

import "iter"

func Cycle[V any](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		var bkp []V
		for v := range seq {
			if !yield(v) {
				return
			}
			bkp = append(bkp, v)
		}
		for {
			for _, v := range bkp {
				if !yield(v) {
					return
				}
			}
		}
	}
}
