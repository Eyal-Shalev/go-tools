package itertools_test

import (
	"github.com/Eyal-Shalev/itertools-go"
	"github.com/Eyal-Shalev/itertools-go/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
)

func TestGroupBy(t *testing.T) {
	type args struct {
		seq   []string
		keyFn func(string) string
	}
	type testCase struct {
		name string
		args args
		want []itertools.Entry[string, []string]
	}
	tests := []testCase{
		{"empty", args{}, nil},
		{
			name: "single_group",
			args: args{[]string{"item", "item", "item"}, internal.Identity[string]},
			want: []itertools.Entry[string, []string]{{"item", []string{"item", "item", "item"}}},
		},
		{
			name: "many_groups",
			args: args{[]string{"item1", "item1", "item2", "item3", "item3"}, internal.Identity[string]},
			want: []itertools.Entry[string, []string]{
				{"item1", []string{"item1", "item1"}},
				{"item2", []string{"item2"}},
				{"item3", []string{"item3", "item3"}},
			},
		},
		{
			name: "unsorted_groups",
			args: args{[]string{"item1", "item2", "item1"}, internal.Identity[string]},
			want: []itertools.Entry[string, []string]{
				{"item1", []string{"item1"}},
				{"item2", []string{"item2"}},
				{"item1", []string{"item1"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := itertools.GroupBy(slices.Values(tt.args.seq), tt.args.keyFn)
			wantOuterIdx := -1
			for key, group := range actual {
				wantOuterIdx += 1
				require.Greaterf(t, len(tt.want), wantOuterIdx, "outer index out of bounds")
				wantCurr := tt.want[wantOuterIdx]
				assert.Equal(t, wantCurr.Key, key)
				wantInnerIdx := -1
				for val := range group {
					wantInnerIdx += 1
					require.Greaterf(t, len(wantCurr.Val), wantInnerIdx, "inner index out of bounds")
					assert.Equal(t, wantCurr.Val[wantInnerIdx], val)
				}
			}
		})
	}
}
