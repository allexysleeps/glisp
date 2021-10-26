package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func argValue(scope *shared.Scope, eval shared.Evaluator, arg shared.ExpArgument) (shared.Value, *shared.Err) {
	switch arg.Type() {
	case shared.ArgTypeValue:
		return arg.(shared.ArgValue).Value, nil
	case shared.ArgTypeVariable:
		vName := arg.(shared.ArgVariable).Value
		vVal, ok := scope.Get(vName)
		if !ok {
			return nil, shared.CreateRootError(shared.ErrUndefined, fmt.Sprintf("undefined variable [%s]", vName), "")
		}
		return vVal.Value(), nil
	case shared.ArgTypeExp:
		return eval(scope, *arg.(shared.ArgExpression).Value)
	}
	return nil, nil
}

func argLenErrorMsg(amount, required int) (string, bool) {
	if amount != required {
		return fmt.Sprintf("invalid amount of argument want %d, got %d", required, amount), false
	}
	return "", true
}
