package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func Print(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) {
	arg := exp.Arguments[0]
	switch arg.Type() {
	case shared.TypeValue:
		fmt.Println(arg.(shared.ArgValue).StrVal())
	case shared.TypeExp:
		fmt.Println(eval(scope, *arg.(shared.ArgExp).Value))
	}
}
