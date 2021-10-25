package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func Def(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) shared.Value {
	if len(exp.Arguments) < 2 {
		panic(fmt.Errorf("invalid ammount of arguments provided to %v: %d want: %d",
			exp.Operation, len(exp.Arguments), 2))
	}

	varName, ok := exp.Arguments[0].(shared.ArgVariable)
	if !ok {
		panic(fmt.Errorf("invalid variable name %v", varName))
	}

	val := argValue(scope, eval, exp.Arguments[1])

	variable := shared.Variable{
		Name:  varName.Value,
		Value: val,
	}

	scope.Set(variable)

	return variable.Value
}
