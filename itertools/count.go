package itertools

import (
	"iter"

	"github.com/Eyal-Shalev/go-tools/mathtools"
	"golang.org/x/exp/constraints"
)

func Count[N constraints.Integer](optFns ...countOption[N]) iter.Seq[N] {
	maxVal := mathtools.MaxN[N]()
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
func WithStart[N mathtools.Rational](start N) countOption[N] {
	return func(options *countOptions[N]) {
		options.start = start
	}
}

//goland:noinspection GoExportedFuncWithUnexportedType
func WithStep[N mathtools.Rational](step N) countOption[N] {
	if step == 0 {
		panic("step cannot be 0")
	}
	return func(options *countOptions[N]) {
		options.step = step
	}
}

type countOption[N mathtools.Rational] func(options *countOptions[N])

type countOptions[N mathtools.Rational] struct {
	start N
	step  N
	end   N
}
