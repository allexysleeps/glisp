package eval

import (
	"fmt"
	"glisp/lang/expression"
)

func Eval(exp expression.Exp) {
	fmt.Printf("%v", exp.Operation)
}