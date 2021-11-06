package shared

import "glisp/lang/errors"

type Evaluator func(scope *Scope, exp Expression) (Value, *errors.Err)
