package itertools_test

import (
	"fmt"
	"testing"

	. "github.com/Eyal-Shalev/go-tools/itertools"
	"github.com/stretchr/testify/assert"
)

func FuzzCount(f *testing.F) {
	f.Fuzz(func(t *testing.T, countStart, countStep, takeEnd int) {
		if countStep == 0 {
			defer func() {
				r := recover()
				assert.NotNilf(t, r, "expected panic on step=0")
			}()
		}
		seq := Take(Count(
			WithStart(countStart),
			WithStep(countStep)),
			WithPredicate(func(v int) bool { return v < takeEnd }))
		for range seq {
		}
	})
}

func ExampleCount() {
	seq := Take(Count(WithStart(-1), WithStep(1)), WithPredicate(func(v int) bool { return v < 2 }))
	for val := range seq {
		fmt.Println(val)
	}
	// Output: -1
	// 0
	// 1
}
