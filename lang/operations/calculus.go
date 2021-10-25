package operations

import (
	"fmt"
	"glisp/lang/shared"
)

func calcArgs(scope *shared.Scope, exp *shared.Exp, eval shared.Evaluator, calc func(a, b float64) float64) shared.Value {
	var r float64
	for i, arg := range exp.Arguments {
		var diff float64
		switch arg.Type() {
		case shared.TypeValue:
			diff = arg.(shared.ArgValue).NumVal()
		case shared.TypeVariable:
			variable := scope.Get(arg.(shared.ArgVariable).Value)
			if variable == nil {
				panic("Nill value")
			}
			diff = variable.NumVal()
		case shared.TypeExp:
			diff = eval(scope, *arg.(shared.ArgExp).Value).NumVal()
		}

		if i == 0 {
			r = diff
		} else {
			r = calc(r, diff)
		}
	}
	res, _ := shared.CreateValue(fmt.Sprintf("%f", r))
	return res
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
