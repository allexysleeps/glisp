package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func Def(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	errMsg, ok := argLenErrorMsg(len(exp.Arguments), 2)
	if !ok {
		return nil, shared.CreateRootError(shared.ErrArgAmount, errMsg, "def")
	}

	varName, ok := exp.Arguments[0].(shared.ArgVariable)
	if !ok {
		return nil, shared.CreateRootError(
			shared.ErrArgAmount, fmt.Sprintf("invalid variable name %v", varName), "def")
	}

	val, err := argValue(scope, eval, exp.Arguments[1])
	if err != nil {
		return nil, shared.CreateErrStack("def", err)
	}

	variable := shared.CreateValueVar(varName.Value, val)

	scope.Set(variable)

	return variable.Value(), nil
}
