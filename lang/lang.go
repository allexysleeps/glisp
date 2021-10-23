package lang

import (
	"fmt"
	"glisp/lang/eval"
	"glisp/lang/shared"
	"io"
)
import (
	"glisp/lang/parse"
)

func Run(input io.Reader) {
	exp := parse.Parse(input)
	scope := shared.CreateScope(nil)
	for _, e := range exp {
		eval.Eval(scope, e)
	}
}

func EvalExp(input io.Reader) {
	exp := parse.Parse(input)[0]
	scope := shared.CreateScope(nil)
	fmt.Printf("glsip>> ")
	fmt.Println(eval.Eval(scope, exp))
}
