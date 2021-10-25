package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func If(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) shared.Value {
	if len(exp.Arguments) < 3 {
		panic(fmt.Errorf("invalid ammount of arguments provided to %v: %d want: %d",
			exp.Operation, len(exp.Arguments), 3))
	}
	cond := argValue(scope, eval, exp.Arguments[0]).BoolVal()
	if cond {
		return argValue(scope, eval, exp.Arguments[1])
	}
	return argValue(scope, eval, exp.Arguments[2])
}

func Eql(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) shared.Value {
	if len(exp.Arguments) < 2 {
		panic(fmt.Errorf("invalid ammount of arguments provided to %v: %d want: %d",
			exp.Operation, len(exp.Arguments), 2))
	}
	v1 := argValue(scope, eval, exp.Arguments[0])
	v2 := argValue(scope, eval, exp.Arguments[1])
	if v1.Type() != v2.Type() {
		return shared.CreateValueOfType(shared.TypeBool, false)
	}
	return shared.CreateValueOfType(shared.TypeBool, v1.StrVal() == v2.StrVal())
}

func More(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) shared.Value {
	return compareNumArgs(scope, exp, eval, func(a, b float64) bool { return a > b })
}

func MoreEq(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) shared.Value {
	return compareNumArgs(scope, exp, eval, func(a, b float64) bool { return a >= b })
}

func compareNumArgs(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator, comp func(a, b float64) bool) shared.Value {
	if len(exp.Arguments) < 2 {
		panic(fmt.Errorf("invalid ammount of arguments provided to %v: %d want: %d",
			exp.Operation, len(exp.Arguments), 2))
	}
	v1 := argValue(scope, eval, exp.Arguments[0]).NumVal()
	v2 := argValue(scope, eval, exp.Arguments[1]).NumVal()
	return shared.CreateValueOfType(shared.TypeBool, comp(v1, v2))
}
