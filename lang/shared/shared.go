package shared

type Evaluator func(scope *Scope, exp Exp) (Value, *Err)
