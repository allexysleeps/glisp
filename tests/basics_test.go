package tests

import (
	"fmt"
	"glisp/lang"
	"strings"
	"testing"
)

func TestEval(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{
			input: "(sum 1 2)",
			want:  "3",
		},
		{
			input: "(sub 10 3)",
			want:  "7",
		},
		{
			input: "(mult 2 2 5)",
			want:  "20",
		},
		{
			input: "(div 125 5)",
			want:  "25",
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
	}
	for _, tst := range tests {
		if res := fmt.Sprintf("%v", lang.EvalExp(strings.NewReader(tst.input))); res != tst.want {
			t.Errorf("lang.EvalExp(%s) == %s, want %s", tst.input, res, tst.want)
		}
	}

}
