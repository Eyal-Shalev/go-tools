package itertools_test

import (
	"iter"
	"math"
	"slices"
	"testing"

	. "github.com/Eyal-Shalev/go-tools/itertools"
	"github.com/Eyal-Shalev/go-tools/mathtools"
	"github.com/stretchr/testify/assert"
)

func FuzzTakeWhile(f *testing.F) {
	f.Fuzz(func(t *testing.T, takeMod int) {
		modFn := mathtools.ModFn(takeMod)
		for range TakeWhile(Count[int](), func(val int) bool { return modFn(val) == 0 }, false) {
		}
	})
}

func FuzzTakeN(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		for range TakeN(Count[int](), n) {
		}
	})
}

func TestTakeN(t *testing.T) {
	type args struct {
		seq iter.Seq[any]
		n   int
	}
	type testCase struct {
		name string
		args args
		want []any
	}
	tests := []testCase{
		{name: "Negative n", args: args{seq: ToAny(Count[int]()), n: -4}, want: nil},
		{name: "Empty Sequence", args: args{seq: ToAny(Count[int]()), n: 0}, want: nil},
		{name: "Sequence [0, 5)", args: args{seq: ToAny(Count[int]()), n: 5}, want: []any{0, 1, 2, 3, 4}},
		{name: "Sequence of strings", args: args{seq: slices.Values([]any{"a", "b", "c"}), n: 2}, want: []any{"a", "b"}},
		{name: "Sequence things", args: args{seq: slices.Values([]any{-4, "foo", 5.4, math.Pi, "boom"}), n: 4}, want: []any{-4, "foo", 5.4, math.Pi}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, slices.Collect(TakeN(tt.args.seq, tt.args.n)))
		})
	}
}
