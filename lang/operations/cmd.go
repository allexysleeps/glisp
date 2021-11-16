package operations

import (
	"fmt"

	"github.com/allexysleeps/glisp/lang/errors"
	"github.com/allexysleeps/glisp/lang/shared"
)

func Print(scope *shared.Scope, exp *shared.Expression, eval shared.Evaluator) (shared.Value, *errors.Err) {
	val, err := argValue(scope, eval, exp.Arguments[0])
	if err != nil {
		return nil, errors.CreateErrStack("print", err)
	}
	fmt.Printf("glisp>> %s\n", val.StrVal())
	return val, nil
}
