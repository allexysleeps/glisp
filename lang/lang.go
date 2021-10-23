package lang

import (
	"glisp/lang/eval"
	"io"
)
import (
	"glisp/lang/parse"
)

func Run(input io.Reader) {
	exp  := parse.Parse(input)
	for _, e := range exp {
		eval.Eval(e)
	}
}