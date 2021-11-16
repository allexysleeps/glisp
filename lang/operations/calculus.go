package operations

import (
	"github.com/allexysleeps/glisp/lang/errors"
	"github.com/allexysleeps/glisp/lang/shared"
)

func calcArgs(
	scope *shared.Scope,
	exp *shared.Expression,
	eval shared.Evaluator,
	operation string,
	calc func(a, b float64) float64) (shared.Value, *errors.Err,
) {
	var r float64
	for i, arg := range exp.Arguments {
		val, err := argValue(scope, eval, arg)
		if err != nil {
			return nil, errors.CreateErrStack(operation, err)
		}
		if i == 0 {
			r = val.NumVal()
		} else {
			r = calc(r, val.NumVal())
		}
	}
	return shared.CreateValueOfType(shared.TypeNum, r), nil
}

func Sum(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *errors.Err) {
	return calcArgs(scope, exp, eval, "sum", func(a, b float64) float64 { return a + b })
}

func Sub(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *errors.Err) {
	return calcArgs(scope, exp, eval, "sub", func(a, b float64) float64 { return a - b })
}

func Mult(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *errors.Err) {
	return calcArgs(scope, exp, eval, "mult", func(a, b float64) float64 { return a * b })
}

func Div(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *errors.Err) {
	return calcArgs(scope, exp, eval, "div", func(a, b float64) float64 { return a / b })
}
