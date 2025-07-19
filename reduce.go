package itertools

import "iter"

func Reduce[V, Acc any](seq iter.Seq[V], reduceFn func(Acc, V) Acc, optFns ...reduceOption[Acc]) Acc {
	var options = reduceOptions[Acc]{}
	for _, optFn := range optFns {
		optFn(&options)
	}
	acc := options.initial
	for v := range seq {
		acc = reduceFn(acc, v)
	}
	return acc
}

func Reduce2[K, V, Acc any](seq iter.Seq2[K, V], reduceFn func(Acc, K, V) Acc, optFns ...reduceOption[Acc]) Acc {
	var options = reduceOptions[Acc]{}
	for _, optFn := range optFns {
		optFn(&options)
	}
	acc := options.initial
	for k, v := range seq {
		acc = reduceFn(acc, k, v)
	}
	return acc
}

type reduceOptions[Acc any] struct {
	initial Acc
}

type reduceOption[Acc any] func(options *reduceOptions[Acc])

//goland:noinspection GoExportedFuncWithUnexportedType
func WithInitial[Acc any](initialAcc Acc) reduceOption[Acc] {
	return func(options *reduceOptions[Acc]) {
		options.initial = initialAcc
	}
}
