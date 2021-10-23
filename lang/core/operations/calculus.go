package operations

import (
	"glisp/lang/expression"
	"glisp/lang/shared"
	"strconv"
)

func Sum(exp *expression.Exp, eval shared.Evaluator) float64 {
	var s float64
	for _, arg := range exp.Arguments {
		switch arg.Type() {
		case expression.TypeValue:
			argVal := arg.(expression.ArgValue)
			numVal, err := strconv.ParseFloat(argVal.Val, 64)
			if err != nil {
				panic(err)
			}
			s += numVal
		case expression.TypeExp:
			s += eval(*arg.(expression.ArgExp).Val).(float64)
		}
	}
	return s
}

func Sub(exp *expression.Exp, eval shared.Evaluator) float64 {
	val := exp.Arguments[0].(expression.ArgValue).Val
	s, _ := strconv.ParseFloat(val, 64)
	for _, arg := range exp.Arguments[1:] {
		switch arg.Type() {
		case expression.TypeValue:
			argVal := arg.(expression.ArgValue)
			numVal, err := strconv.ParseFloat(argVal.Val, 64)
			if err != nil {
				panic(err)
			}
			s -= numVal
		case expression.TypeExp:
			s -= eval(*arg.(expression.ArgExp).Val).(float64)
		}
	}
	return s
}
