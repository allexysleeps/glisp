package eval

import (
	"glisp/lang/errors"
	"glisp/lang/operations"
	"glisp/lang/shared"
)

type operation = func(s *shared.Scope, e *shared.Expression, eval shared.Evaluator) (shared.Value, *errors.Err)

var operationsMap = map[string]operation{
	// calculus
	"sum":  operations.Sum,
	"sub":  operations.Sub,
	"mult": operations.Mult,
	"div":  operations.Div,
	// logical
	"if":     operations.If,
	"eql":    operations.Eql,
	"more":   operations.More,
	"moreEq": operations.MoreEq,
	// define
	"def": operations.Def,
	"fn":  operations.DefFn,
	// list
	"list":    operations.List,
	"length":  operations.Length,
	"map":     operations.Map,
	"get":     operations.Get,
	"sublist": operations.SubList,
	// cmd
	"print": operations.Print,
}

func Eval(parentScope *shared.Scope, exp shared.Expression) shared.Value {
	val, err := eval(parentScope, exp)
	if err != nil {
		err.Print()
		return nil
	}
	return val
}

func eval(scope *shared.Scope, exp shared.Expression) (shared.Value, *errors.Err) {
	op, ok := operationsMap[exp.Operation]
	if !ok {
		return operations.Function(exp.Operation, scope, exp, eval)
	}
	val, err := op(scope, &exp, eval)
	if err != nil {
		return nil, err
	}
	return val, nil
}
