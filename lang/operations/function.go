package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func Function(fName string, parentScope *shared.Scope, parentExp shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	val, ok := parentScope.Get(fName)
	if !ok {
		return nil, shared.CreateRootError(shared.ErrUndefined, fmt.Sprintf("undefined function name %scope", fName), "")
	}
	if val.Type() != shared.VarFn {
		return nil, shared.CreateRootError(shared.ErrUndefined, fmt.Sprintf("%scope is not a function", fName), "")
	}
	exp, fArgs := val.Exec()
	scope := shared.CreateScope(parentScope)
	if len(fArgs) != len(parentExp.Arguments) {
		return nil, shared.CreateRootError(shared.ErrArgAmount, fmt.Sprintf(
			"wrong amount of arguments provided to %scope want %d, got %d", fName, len(fArgs), len(parentExp.Arguments)),
			fName)
	}
	for i, argName := range fArgs {
		argVal, err := argValue(parentScope, eval, parentExp.Arguments[i])
		if err != nil {
			return nil, shared.CreateErrStack(fName, err)
		}
		scope.Set(shared.CreateValueVar(argName, argVal))
	}
	return eval(scope, *exp)
}
