package itertools

import (
	"iter"

	"github.com/Eyal-Shalev/go-tools/mathtools"
	"golang.org/x/exp/constraints"
)

func Count[N constraints.Integer]() iter.Seq[N] {
	maxVal := mathtools.MaxN[N]()
	return func(yield func(N) bool) {
		for i := N(0); i < maxVal; i++ {
			if !yield(i) {
				return
			}
		}
	}
}
