package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func List(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	var values []shared.Value

	for _, arg := range exp.Arguments {
		val, err := argValue(scope, eval, arg)
		if err != nil {
			return nil, shared.CreateErrStack("list", err)
		}
		values = append(values, val)
	}

	val := shared.CreateValueOfType(shared.TypeList, &values)

	return val, nil
}

func Length(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	errMsg, ok := argLenErrorMsg(len(exp.Arguments), 1)
	if !ok {
		return nil, shared.CreateRootError(shared.ErrArgAmount, errMsg, "length")
	}

	list, err := argValue(scope, eval, exp.Arguments[0])
	if err != nil {
		return nil, shared.CreateErrStack("length", err)
	}
	if list.Type() != shared.TypeList {
		return nil, shared.CreateRootError(shared.ErrWrongSyntax, "arg provided at 1 to get is not of a list type", "length")
	}

	return shared.CreateValueOfType(shared.TypeNum, float64(len(*list.ListVal()))), nil
}

func Get(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	errMsg, ok := argLenErrorMsg(len(exp.Arguments), 2)
	if !ok {
		return nil, shared.CreateRootError(shared.ErrArgAmount, errMsg, "get")
	}

	index, err := argValue(scope, eval, exp.Arguments[0])
	if err != nil {
		return nil, shared.CreateErrStack("map", err)
	}
	if index.Type() != shared.TypeNum {
		return nil, shared.CreateRootError(shared.ErrWrongSyntax, "arg provided at 1 to get is not of a num type", "get")
	}

	list, err := argValue(scope, eval, exp.Arguments[1])
	if err != nil {
		return nil, shared.CreateErrStack("map", err)
	}
	if list.Type() != shared.TypeList {
		return nil, shared.CreateRootError(shared.ErrWrongSyntax, "arg provided at 2 to get is not of a list type", "get")
	}

	listValues := list.ListVal()

	return (*listValues)[int(index.NumVal())], nil
}

func SubList(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *shared.Err) {
	argsLen := len(exp.Arguments)
	if argsLen < 1 {
		return nil, shared.CreateRootError(shared.ErrArgAmount, "to few arguments provided to subList", "subList")
	}

	list, err := argValue(scope, eval, exp.Arguments[0])
	if err != nil {
		return nil, shared.CreateErrStack("map", err)
	}
	if list.Type() != shared.TypeList {
		return nil, shared.CreateRootError(shared.ErrWrongSyntax, "arg provided at 1 to get is not of a list type", "subList")
	}

	start, err := argValue(scope, eval, exp.Arguments[1])
	if err != nil {
		return nil, shared.CreateErrStack("map", err)
	}
	if start.Type() != shared.TypeNum {
		return nil, shared.CreateRootError(shared.ErrWrongSyntax, "start provided to subList is not of a num type", "subList")
	}

	var end int
	if argsLen == 3 {
		endVal, err := argValue(scope, eval, exp.Arguments[2])
		if err != nil {
			return nil, shared.CreateErrStack("map", err)
		}
		if endVal.Type() != shared.TypeNum {
			return nil, shared.CreateRootError(shared.ErrWrongSyntax, "end provided to subList is not of a num type", "subList")
		}
		if end > argsLen-1 {
			return nil, shared.CreateRootError(shared.ErrWrongSyntax, "end provided to subList is out of list range", "subList")
		}
		end = int(endVal.NumVal())
	} else {
		end = argsLen - 1
	}

	newList := (*list.ListVal())[int(start.NumVal()) : end+1]

	return shared.CreateValueOfType(shared.TypeList, &newList), nil
}

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
