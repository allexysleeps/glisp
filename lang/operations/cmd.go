package operations

import (
	"log"

	"glisp/lang/errors"
	"glisp/lang/shared"
)

func Print(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *errors.Err) {
	val, err := argValue(scope, eval, exp.Arguments[0])
	if err != nil {
		return nil, errors.CreateErrStack("print", err)
	}
	log.Print(val.StrVal())
	return val, nil
}
