package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func Print(exp *shared.Exp, eval shared.Evaluator) {
	arg := exp.Arguments[0]
	switch arg.Type() {
	case shared.TypeValue:
		fmt.Println(arg.(shared.ArgValue).Value)
	case shared.TypeExp:
		fmt.Println(eval(*arg.(shared.ArgExp).Value))
	}
}
