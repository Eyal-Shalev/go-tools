package itertools

import "iter"

func GroupBy[K comparable, V any](seq iter.Seq[V], keyFn func(V) K) iter.Seq2[K, iter.Seq[V]] {
	return func(outerYield func(K, iter.Seq[V]) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		v, ok := next()
		for ok {
			lastKey := keyFn(v)

			var innerIter iter.Seq[V] = func(innerYield func(V) bool) {
				for ; keyFn(v) == lastKey && ok; v, ok = next() {
					if !innerYield(v) {
						return
					}
				}
			}

			if !outerYield(lastKey, innerIter) {
				return
			}
		}
	}
}
