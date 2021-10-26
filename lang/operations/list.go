package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func Map(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	errMsg, ok := argLenErrorMsg(len(exp.Arguments), 2)
	if !ok {
		return nil, shared.CreateRootError(shared.ErrArgAmount, errMsg, "map")
	}

	fName := exp.Arguments[0].(shared.ArgVariable).Value
	_, ok = scope.Get(fName)
	if !ok {
		return nil, shared.CreateRootError(shared.ErrUndefined, fmt.Sprintf("undefined function %s", fName), "map")
	}

	list, err := argValue(scope, eval, exp.Arguments[1])
	if err != nil {
		return nil, shared.CreateErrStack("map", err)
	}
	if list.Type() != shared.TypeList {
		return nil, shared.CreateRootError(shared.ErrWrongSyntax, "arg provided to map is not of a list type", "map")
	}

	var values []shared.Value

	for _, li := range *list.ListVal() {
		val, err := execFunction(fName, scope, eval, []shared.Value{li})
		if err != nil {
			return nil, shared.CreateErrStack("map", err)
		}
		values = append(values, val)
	}

	return shared.CreateValueOfType(shared.TypeList, &values), nil
}
