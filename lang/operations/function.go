package operations

import (
	"fmt"

	"glisp/lang/errors"
	"glisp/lang/shared"
)

func DefFn(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *errors.Err) {
	if exp.Arguments[0].Type() != shared.ArgTypeVariable {
		return nil, errors.CreateRootError(
			errors.ErrWrongSyntax, fmt.Sprintf("incorrect function name %v", exp.Arguments[0]), "fn")
	}
	fName := exp.Arguments[0].(shared.ArgVariable).Value
	var arity int
	var fArgs []string
	for i, arg := range exp.Arguments[1:] {
		if arg.Type() != shared.ArgTypeArgument {
			arity = i
			break
		}
		fArgs = append(fArgs, arg.(shared.ArgArgument).Value)
	}

	if len(exp.Arguments) != arity+2 {
		return nil, errors.CreateRootError(errors.ErrWrongSyntax, fmt.Sprintf(
			"incorrect function syntax for %s, incorrect arg amount", fName), "fn")
	}

	expArg := exp.Arguments[arity+1]
	if expArg.Type() != shared.ArgTypeExp {
		return nil, errors.CreateRootError(errors.ErrWrongSyntax, fmt.Sprintf(
			"incorrect function syntax for %s", fName), "fn")
	}

	fExp := expArg.(shared.ArgExpression).Value

	fVar := shared.CreateFunctionVar(fName, fExp, fArgs)
	scope.Set(fVar)

	return nil, nil
}

func Function(fName string, parentScope *shared.Scope,
	parentExp shared.Expression, eval shared.Evaluator) (shared.Value, *errors.Err) {
	val, ok := parentScope.Get(fName)
	if !ok {
		return nil, errors.CreateRootError(
			errors.ErrUndefined, fmt.Sprintf("undefined function name %scope", fName), "")
	}
	if val.Type() != shared.VarFn {
		return nil, errors.CreateRootError(
			errors.ErrUndefined, fmt.Sprintf("%scope is not a function", fName), "")
	}
	exp, fArgs := val.Exec()
	scope := shared.CreateScope(parentScope)
	if len(fArgs) != len(parentExp.Arguments) {
		return nil, errors.CreateRootError(errors.ErrArgAmount, fmt.Sprintf(
			"wrong amount of arguments provided to %scope want %d, got %d", fName, len(fArgs), len(parentExp.Arguments)),
			fName)
	}
	for i, argName := range fArgs {
		argVal, err := argValue(parentScope, eval, parentExp.Arguments[i])
		if err != nil {
			return nil, errors.CreateErrStack(fName, err)
		}
		scope.Set(shared.CreateValueVar(argName, argVal))
	}
	return eval(scope, *exp)
}

func execFunction(fName string, parentScope *shared.Scope,
	eval shared.Evaluator, argValues []shared.Value) (shared.Value, *errors.Err) {
	val, ok := parentScope.Get(fName)
	if !ok {
		return nil, errors.CreateRootError(
			errors.ErrUndefined, fmt.Sprintf("undefined function name %scope", fName), "")
	}
	if val.Type() != shared.VarFn {
		return nil, errors.CreateRootError(
			errors.ErrUndefined, fmt.Sprintf("%scope is not a function", fName), "")
	}
	exp, fArgs := val.Exec()
	scope := shared.CreateScope(parentScope)
	if len(fArgs) != len(argValues) {
		return nil, errors.CreateRootError(errors.ErrArgAmount, fmt.Sprintf(
			"wrong amount of arguments provided to %scope want %d, got %d", fName, len(fArgs), len(argValues)),
			fName)
	}
	for i, argName := range fArgs {
		argVal := argValues[i]
		scope.Set(shared.CreateValueVar(argName, argVal))
	}
	return eval(scope, *exp)
}
