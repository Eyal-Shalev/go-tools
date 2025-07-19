package itertools_test

import (
	. "github.com/Eyal-Shalev/itertools-go"
	. "github.com/Eyal-Shalev/itertools-go/internal"
	"testing"
)

func FuzzTake(f *testing.F) {
	f.Fuzz(func(t *testing.T, countStart, countStep, countEnd int, takeMod int) {
		modFn := ModFn(takeMod)
		Take(Count(WithStart(countStart), WithStep(countStep)), WithPredicate[int](func(val int) bool {
			return modFn(val) == 0
		}))
	})
}
