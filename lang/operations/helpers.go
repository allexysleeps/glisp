package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func argValue(scope *shared.Scope, eval shared.Evaluator, arg shared.ExpArgument) (shared.Value, *shared.Err) {
	var val shared.Value
	var err *shared.Err
	switch arg.Type() {
	case shared.TypeValue:
		val = arg.(shared.ArgValue).Value
	case shared.TypeVariable:
		vname := arg.(shared.ArgVariable).Value
		vval, ok := scope.Get(vname)
		val = vval
		if !ok {
			err = shared.CreateRootError(shared.ErrUndefined, fmt.Sprintf("undefined variable [%s]", vname), "")
		}
	case shared.TypeExp:
		val, err = eval(scope, *arg.(shared.ArgExp).Value)
	}
	return val, err
}

func argLenErrorMsg(amount, required int) (string, bool) {
	if amount != required {
		return fmt.Sprintf("invalid amount of argument want %d, got %d", required, amount), false
	}
	return "", true
}
