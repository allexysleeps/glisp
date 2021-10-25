package operations

import (
	"glisp/lang/shared"
)

func calcArgs(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator, calc func(a, b float64) float64) shared.Value {
	var r float64
	for i, arg := range exp.Arguments {
		val := argValue(scope, eval, arg).NumVal()
		if i == 0 {
			r = val
		} else {
			r = calc(r, val)
		}
	}
	return shared.CreateValueOfType(shared.TypeNum, r)
}

func Sum(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) shared.Value {
	return calcArgs(scope, exp, eval, func(a, b float64) float64 { return a + b })
}

func Sub(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) shared.Value {
	return calcArgs(scope, exp, eval, func(a, b float64) float64 { return a - b })
}

func Mult(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) shared.Value {
	return calcArgs(scope, exp, eval, func(a, b float64) float64 { return a * b })
}

func Div(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) shared.Value {
	return calcArgs(scope, exp, eval, func(a, b float64) float64 { return a / b })
}
