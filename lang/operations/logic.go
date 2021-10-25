package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func If(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) shared.Value {
	if len(exp.Arguments) < 3 {
		panic(fmt.Errorf("invalid ammount of arguments provided to %v: %d want: %d",
			exp.Operation, len(exp.Arguments), 2))
	}
	cond := argValue(scope, eval, exp.Arguments[0]).BoolVal()
	if cond {
		return argValue(scope, eval, exp.Arguments[1])
	}
	return argValue(scope, eval, exp.Arguments[2])
}
