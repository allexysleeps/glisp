package eval

import (
	"glisp/lang/core/operations"
	"glisp/lang/expression"
	"log"
)

func Eval(exp expression.Exp) interface{} {
	return eval(exp)
}

func eval(exp expression.Exp) interface{} {
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
	default:
		log.Printf("undefined operation %v\n", exp.Operation)
	}
	return nil
}
