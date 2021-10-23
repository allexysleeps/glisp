package shared

import "glisp/lang/expression"

type Evaluator func(exp expression.Exp) interface{}
