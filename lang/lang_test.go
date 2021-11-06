package lang

import (
	"fmt"
	"strings"
	"testing"
)

func TestEval(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		input string
		want  string
	}{
		{
			input: "(sum 1 2)",
			want:  "3",
		},
		{
			input: "(mult 2 2 5)",
			want:  "20",
		},
		{
			input: "(sum (div 100 2 (sum 2 3)) (sub 20 7 3))",
			want:  "20",
		},
		{
			input: "(if true 10 20)",
			want:  "10",
		},
		{
			input: "(if false 10 20)",
			want:  "20",
		},
		{
			input: "(eql 5 5)",
			want:  "true",
		},
		{
			input: "(eql 5 6)",
			want:  "false",
		},
		{
			input: "(more 7 6)",
			want:  "true",
		},
		{
			input: "(more 5 6)",
			want:  "false",
		},
		{
			input: "(moreEq 6 6)",
			want:  "true",
		},
		{
			input: "(moreEq 9 6)",
			want:  "true",
		},
	}
	for _, tst := range tests {
		if res := fmt.Sprintf("%v", EvalExp(strings.NewReader(tst.input))); res != tst.want {
			t.Errorf("lang.EvalExp(%s) == %s, want %s", tst.input, res, tst.want)
		}
	}
}
