package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func Def(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) *shared.Variable {
	if len(exp.Arguments) < 2 {
		panic(fmt.Errorf("invalid ammount of arguments provided to %v: %d want: %d",
			exp.Operation, len(exp.Arguments), 2))
	}

	varName, ok := exp.Arguments[0].(shared.ArgVariable)
	if !ok {
		panic(fmt.Errorf("invalid variable name %v", varName))
	}

	var val string
	switch exp.Arguments[1].Type() {
	case shared.TypeValue:
		val = exp.Arguments[1].(shared.ArgValue).StrVal()
	case shared.TypeExp:
		val = fmt.Sprintf("%v", eval(scope, *exp.Arguments[1].(shared.ArgExp).Value))
	}
	variable := shared.Variable{
		Name:  varName.Value,
		Value: val,
	}

	scope.Set(&variable)

	return &variable
}