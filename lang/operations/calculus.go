package operations

import (
	"glisp/lang/shared"
)

func calcArgs(exp *shared.Exp, eval shared.Evaluator, calc func(a, b float64) float64) float64 {
	var r float64
	for i, arg := range exp.Arguments {
		var diff float64

		switch arg.Type() {
		case shared.TypeValue:
			diff = getNumArg(arg)
		case shared.TypeExp:
			diff = eval(*arg.(shared.ArgExp).Value).(float64)
		}

		if i == 0 {
			r = diff
		} else {
			r = calc(r, diff)
		}
	}
	return r
}

func Sum(exp *shared.Exp, eval shared.Evaluator) float64 {
	return calcArgs(exp, eval, func(a, b float64) float64 { return a + b })
}

func Sub(exp *shared.Exp, eval shared.Evaluator) float64 {
	return calcArgs(exp, eval, func(a, b float64) float64 { return a - b })
}

func Mult(exp *shared.Exp, eval shared.Evaluator) float64 {
	return calcArgs(exp, eval, func(a, b float64) float64 { return a * b })
}

func Div(exp *shared.Exp, eval shared.Evaluator) float64 {
	return calcArgs(exp, eval, func(a, b float64) float64 { return a / b })
}
