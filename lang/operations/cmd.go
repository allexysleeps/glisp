package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func Print(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	val, err := argValue(scope, eval, exp.Arguments[0])
	if err != nil {
		return nil, shared.CreateErrStack("print", err)
	}
	fmt.Println(val.StrVal())
	return val, nil
}
