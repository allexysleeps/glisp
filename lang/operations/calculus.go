package operations

import (
	"glisp/lang/shared"
)

func calcArgs(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator, calc func(a, b float64) float64) float64 {
	var r float64
	for i, arg := range exp.Arguments {
		var diff float64
		switch arg.Type() {
		case shared.TypeValue:
			diff = getNumArg(arg)
		case shared.TypeVariable:
			variable := scope.Get(arg.(shared.ArgVariable).Value)
			val, err := variable.NumVal()
			if err != nil {
				panic(err)
			}
			diff = val
		case shared.TypeExp:
			diff = eval(scope, *arg.(shared.ArgExp).Value).(float64)
		}

		if i == 0 {
			r = diff
		} else {
			r = calc(r, diff)
		}
	}
	return r
}

func Sum(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) float64 {
	return calcArgs(scope, exp, eval, func(a, b float64) float64 { return a + b })
}

func Sub(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) float64 {
	return calcArgs(scope, exp, eval, func(a, b float64) float64 { return a - b })
}

func Mult(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) float64 {
	return calcArgs(scope, exp, eval, func(a, b float64) float64 { return a * b })
}

func Div(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator) float64 {
	return calcArgs(scope, exp, eval, func(a, b float64) float64 { return a / b })
}
