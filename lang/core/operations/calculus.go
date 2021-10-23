package operations

import (
	"glisp/lang/expression"
	"glisp/lang/shared"
)

func calcArgs(exp *expression.Exp, eval shared.Evaluator, calc func(a, b float64) float64) float64 {
	var r float64
	for i, arg := range exp.Arguments {
		var diff float64

		switch arg.Type() {
		case expression.TypeValue:
			diff = getNumArg(arg)
		case expression.TypeExp:
			diff = eval(*arg.(expression.ArgExp).Val).(float64)
		}

		if i == 0 {
			r = diff
		} else {
			r = calc(r, diff)
		}
	}
	return r
}

func Sum(exp *expression.Exp, eval shared.Evaluator) float64 {
	return calcArgs(exp, eval, func(a, b float64) float64 { return a + b })
}

func Sub(exp *expression.Exp, eval shared.Evaluator) float64 {
	return calcArgs(exp, eval, func(a, b float64) float64 { return a - b })
}

func Mult(exp *expression.Exp, eval shared.Evaluator) float64 {
	return calcArgs(exp, eval, func(a, b float64) float64 { return a * b })
}

func Div(exp *expression.Exp, eval shared.Evaluator) float64 {
	return calcArgs(exp, eval, func(a, b float64) float64 { return a / b })
}
