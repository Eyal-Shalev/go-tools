package itertools

import (
	"github.com/Eyal-Shalev/itertools-go/internal"
	"golang.org/x/exp/constraints"
	"iter"
)

func Count[N constraints.Integer](optFns ...countOption[N]) iter.Seq[N] {
	maxVal := internal.MaxN[N]()
	options := countOptions[N]{
		start: 0,
		step:  1,
	}
	for _, optFn := range optFns {
		optFn(&options)
	}
	val := options.start
	return func(yield func(N) bool) {
		for {
			if !yield(val) {
				return
			}
			if val > maxVal-options.step {
				return
			}
			val += options.step
		}
	}
}

//goland:noinspection GoExportedFuncWithUnexportedType
func WithStart[N internal.Rational](start N) countOption[N] {
	return func(options *countOptions[N]) {
		options.start = start
	}
}

//goland:noinspection GoExportedFuncWithUnexportedType
func WithStep[N internal.Rational](step N) countOption[N] {
	if step == 0 {
		panic("step cannot be 0")
	}
	return func(options *countOptions[N]) {
		options.step = step
	}
}

type countOption[N internal.Rational] func(options *countOptions[N])

type countOptions[N internal.Rational] struct {
	start N
	step  N
	end   N
}
