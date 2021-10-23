package lang

import (
	"fmt"
	"glisp/lang/eval"
	"io"
)
import (
	"glisp/lang/parse"
)

func Run(input io.Reader) {
	exp := parse.Parse(input)
	for _, e := range exp {
		eval.Eval(e)
	}
}

func EvalExp(input io.Reader) {
	exp := parse.Parse(input)[0]
	fmt.Printf("glsip>> ")
	fmt.Println(eval.Eval(exp))
}
