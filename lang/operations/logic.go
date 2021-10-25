package operations

import (
	"glisp/lang/shared"
)

func If(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	errMsg, ok := argLenErrorMsg(len(exp.Arguments), 3)
	if !ok {
		return nil, shared.CreateRootError(shared.ErrArgAmount, errMsg, "if")
	}
	cond, err := argValue(scope, eval, exp.Arguments[0])
	if err != nil {
		shared.CreateErrStack("if", err)
	}
	if cond.BoolVal() {
		val, err := argValue(scope, eval, exp.Arguments[1])
		if err != nil {
			return nil, shared.CreateErrStack("if", err)
		}
		return val, nil
	}
	val, err := argValue(scope, eval, exp.Arguments[2])
	if err != nil {
		return nil, shared.CreateErrStack("if", err)
	}
	return val, nil
}

func Eql(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	errMsg, ok := argLenErrorMsg(len(exp.Arguments), 2)
	if !ok {
		return nil, shared.CreateRootError(shared.ErrArgAmount, errMsg, "eql")
	}
	v1, err := argValue(scope, eval, exp.Arguments[0])
	if err != nil {
		return nil, shared.CreateErrStack("eql", err)
	}
	v2, err := argValue(scope, eval, exp.Arguments[1])
	if err != nil {
		return nil, shared.CreateErrStack("eql", err)
	}
	if v1.Type() != v2.Type() {
		return shared.CreateValueOfType(shared.TypeBool, false), nil
	}
	return shared.CreateValueOfType(shared.TypeBool, v1.StrVal() == v2.StrVal()), nil
}

func More(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	return compareNumArgs(scope, exp, eval, "more", func(a, b float64) bool { return a > b })
}

func MoreEq(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	return compareNumArgs(scope, exp, eval, "moreEq", func(a, b float64) bool { return a >= b })
}

func compareNumArgs(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator, operation string, comp func(a, b float64) bool) (shared.Value, *shared.Err) {
	errMsg, ok := argLenErrorMsg(len(exp.Arguments), 2)
	if !ok {
		return nil, shared.CreateRootError(shared.ErrArgAmount, errMsg, operation)
	}

	v1, err := argValue(scope, eval, exp.Arguments[0])
	if err != nil {
		return nil, shared.CreateErrStack(operation, err)
	}
	v2, err := argValue(scope, eval, exp.Arguments[1])
	if err != nil {
		return nil, shared.CreateErrStack(operation, err)
	}
	return shared.CreateValueOfType(shared.TypeBool, comp(v1.NumVal(), v2.NumVal())), nil
}
