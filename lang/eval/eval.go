package eval

import (
	"glisp/lang/operations"
	"glisp/lang/shared"
	"log"
)

func Eval(parentScope *shared.Scope, exp shared.Exp) shared.Value {
	return eval(parentScope, exp)
}

func eval(scope *shared.Scope, exp shared.Exp) shared.Value {
	switch exp.Operation {
	case "sum":
		return operations.Sum(scope, &exp, eval)
	case "sub":
		return operations.Sub(scope, &exp, eval)
	case "mult":
		return operations.Mult(scope, &exp, eval)
	case "div":
		return operations.Div(scope, &exp, eval)
	case "print":
		operations.Print(scope, &exp, eval)
		return nil
	case "def":
		return operations.Def(scope, &exp, eval)
	default:
		log.Printf("undefined operation %v\n", exp.Operation)
	}
	return nil
}
