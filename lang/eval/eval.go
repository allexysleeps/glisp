package eval

import (
	"fmt"
	"glisp/lang/operations"
	"glisp/lang/shared"
	"log"
)

type operation = func(s *shared.Scope, e *shared.Exp, eval shared.Evaluator) (shared.Value, *shared.Err)

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
	//cmd
	"print": operations.Print,
}

func Eval(parentScope *shared.Scope, exp shared.Exp) shared.Value {
	val, err := eval(parentScope, exp)
	if err != nil {
		err.Print()
		return nil
	}
	return val
}

func eval(scope *shared.Scope, exp shared.Exp) (shared.Value, *shared.Err) {
	op, ok := operationsMap[exp.Operation]
	if !ok {
		log.Printf("undefined operation %v\n", exp.Operation)
		return nil, shared.CreateRootError(shared.ErrUndefined, fmt.Sprintf("undefined operation"), exp.Operation)
	}
	val, err := op(scope, &exp, eval)
	if err != nil {
		return nil, err
	}
	return val, nil
}
