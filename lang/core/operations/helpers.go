package operations

import (
	"glisp/lang/expression"
	"strconv"
)

func getNumArg(val expression.ExpArgument) float64 {
	argVal := val.(expression.ArgValue)
	numVal, err := strconv.ParseFloat(argVal.Val, 64)
	if err != nil {
		panic(err)
	}
	return numVal
}
