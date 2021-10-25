package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func Print(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) shared.Value {
	val := argValue(scope, eval, exp.Arguments[0])
	fmt.Println(val.StrVal())
	return val
}
