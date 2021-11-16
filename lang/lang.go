package lang

import (
	"io"

	"github.com/allexysleeps/glisp/lang/eval"
	"github.com/allexysleeps/glisp/lang/parse"
	"github.com/allexysleeps/glisp/lang/shared"
)

func Run(input io.Reader) {
	exp := parse.Parse(input)
	scope := shared.CreateScope(nil)
	for _, e := range exp {
		eval.Eval(scope, e)
	}
}

func EvalExp(input io.Reader) interface{} {
	exp := parse.Parse(input)[0]
	scope := shared.CreateScope(nil)
	res := eval.Eval(scope, exp)
	return res.StrVal()
}
