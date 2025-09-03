package itertools

import "iter"

func SelectMap[V, VOut any](seq iter.Seq[V], f func(V) (VOut, bool)) iter.Seq[VOut] {
	return func(yield func(VOut) bool) {
		for v := range seq {
			out, ok := f(v)
			if !ok {
				continue
			}
			if !yield(out) {
				return
			}
		}
	}
}

func SelectMap2[K, V, KOut, VOut any](seq iter.Seq2[K, V], f func(K, V) (KOut, VOut, bool)) iter.Seq2[KOut, VOut] {
	return func(yield func(KOut, VOut) bool) {
		for k, v := range seq {
			kOut, vOut, ok := f(k, v)
			if !ok {
				continue
			}
			if !yield(kOut, vOut) {
				return
			}
		}
	}
}
