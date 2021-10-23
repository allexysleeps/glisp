package operations

import (
	"glisp/lang/shared"
	"strconv"
)

func getNumArg(val shared.ExpArgument) float64 {
	argVal := val.(shared.ArgValue)
	numVal, err := strconv.ParseFloat(argVal.Value, 64)
	if err != nil {
		panic(err)
	}
	return numVal
}
