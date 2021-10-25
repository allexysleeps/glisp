package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func Def(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	errMsg, ok := argLenErrorMsg(len(exp.Arguments), 2)
	if !ok {
		return nil, shared.CreateRootError(shared.ErrArgAmount, errMsg, "dev")
	}

	varName, ok := exp.Arguments[0].(shared.ArgVariable)
	if !ok {
		panic(fmt.Errorf("invalid variable name %v", varName))
	}

	val, err := argValue(scope, eval, exp.Arguments[1])
	if err != nil {
		return nil, shared.CreateErrStack("def", err)
	}

	variable := shared.CreateValueVar(varName.Value, val)

	scope.Set(variable)

	return variable.Value(), nil
}

func DefFn(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	if exp.Arguments[0].Type() != shared.TypeVariable {
		return nil, shared.CreateRootError(shared.ErrWrongSyntax, fmt.Sprintf("incorrect function name %v", exp.Arguments[0]), "fn")
	}
	fName := exp.Arguments[0].(shared.ArgVariable).Value
	var arity int
	var fArgs []string
	for i, arg := range exp.Arguments[1:] {
		if arg.Type() != shared.TypeArgument {
			arity = i
			break
		}
		fArgs = append(fArgs, arg.(shared.ArgArgument).Value)
	}

	if len(exp.Arguments) != arity+2 {
		return nil, shared.CreateRootError(shared.ErrWrongSyntax, fmt.Sprintf("incorrect function syntax for %s, incorrect arg amount", fName), "fn")
	}

	expArg := exp.Arguments[arity+1]
	if expArg.Type() != shared.TypeExp {
		return nil, shared.CreateRootError(shared.ErrWrongSyntax, fmt.Sprintf("incorrect function syntax for %s", fName), "fn")
	}

	fExp := expArg.(shared.ArgExpression).Value
	fScope := shared.CreateScope(scope)

	fVar := shared.CreateFunctionVar(fName, fExp, *fScope, fArgs)
	scope.Set(fVar)

	return nil, nil
}
