package itertools_test

import (
	"fmt"
	"testing"

	. "github.com/Eyal-Shalev/go-tools/itertools"
)

func FuzzCount(f *testing.F) {
	f.Fuzz(func(t *testing.T, takeEnd int) {
		seq := TakeWhile(Count[int](), func(v int) bool { return v < takeEnd }, false)
		for range seq {
		}
	})
}

func ExampleCount() {
	seq := TakeWhile(Count[int](), func(v int) bool { return v < 2 }, false)
	for val := range seq {
		fmt.Println(val)
	}
	// Output: 0
	// 1
}
