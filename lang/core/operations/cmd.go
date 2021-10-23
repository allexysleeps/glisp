package operations

import (
	"fmt"
	"glisp/lang/expression"
	"glisp/lang/shared"
)

func Print(exp *expression.Exp, eval shared.Evaluator) {
	arg := exp.Arguments[0]
	switch arg.Type() {
	case expression.TypeValue:
		fmt.Println(arg.(expression.ArgValue).Val)
	case expression.TypeExp:
		fmt.Println(eval(*arg.(expression.ArgExp).Val))
	}
}
