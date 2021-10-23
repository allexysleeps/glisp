package eval

import (
	"glisp/lang/operations"
	"glisp/lang/shared"
	"glisp/lang/variables"
	"log"
)

func Eval(exp shared.Exp) interface{} {
	return eval(exp)
}

func eval(exp shared.Exp) interface{} {
	scope := new(shared.Scope)
	scope.Create()
	exp.Scope = scope
	switch exp.Operation {
	case "sum":
		return operations.Sum(&exp, eval)
	case "sub":
		return operations.Sub(&exp, eval)
	case "mult":
		return operations.Mult(&exp, eval)
	case "div":
		return operations.Div(&exp, eval)
	case "print":
		operations.Print(&exp, eval)
		return nil
	case "def":
		return variables.Def(&exp, eval)
	default:
		log.Printf("undefined operation %v\n", exp.Operation)
	}
	return nil
}
