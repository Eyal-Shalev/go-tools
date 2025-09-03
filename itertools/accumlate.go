package itertools

import "iter"

func Accumulate[V, Acc any](seq iter.Seq[V], accumulateFn func(Acc, V) Acc, optFns ...reduceOption[Acc]) iter.Seq[Acc] {
	var options = reduceOptions[Acc]{}
	for _, optFn := range optFns {
		optFn(&options)
	}
	acc := options.initial
	return func(yield func(Acc) bool) {
		for v := range seq {
			acc = accumulateFn(acc, v)
			if !yield(acc) {
				return
			}
		}
	}
}
