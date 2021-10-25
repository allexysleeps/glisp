package eval

import (
	"glisp/lang/operations"
	"glisp/lang/shared"
	"log"
)

type operation = func(s *shared.Scope, e *shared.Exp, eval shared.Evaluator) shared.Value

var operationsMap = map[string]operation{
	// calculus
	"sum":  operations.Sum,
	"sub":  operations.Sub,
	"mult": operations.Mult,
	"div":  operations.Div,
	// logical
	"if": operations.If,
	// define
	"def": operations.Def,
	//cmd
	"print": operations.Print,
}

func Eval(parentScope *shared.Scope, exp shared.Exp) shared.Value {
	return eval(parentScope, exp)
}

func eval(scope *shared.Scope, exp shared.Exp) shared.Value {
	op, ok := operationsMap[exp.Operation]
	if !ok {
		log.Printf("undefined operation %v\n", exp.Operation)
		return nil
	}
	return op(scope, &exp, eval)
}
