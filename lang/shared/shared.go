package shared

import (
	"github.com/allexysleeps/glisp/lang/errors"
)

type Evaluator func(scope *Scope, exp Expression) (Value, *errors.Err)
