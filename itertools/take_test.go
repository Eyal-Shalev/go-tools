package itertools_test

import (
	"testing"

	. "github.com/Eyal-Shalev/go-tools/itertools"
	"github.com/Eyal-Shalev/go-tools/mathtools"
)

func FuzzTake(f *testing.F) {
	f.Fuzz(func(t *testing.T, countStart, countStep, countEnd int, takeMod int) {
		modFn := mathtools.ModFn(takeMod)
		Take(Count(WithStart(countStart), WithStep(countStep)), WithPredicate[int](func(val int) bool {
			return modFn(val) == 0
		}))
	})
}
