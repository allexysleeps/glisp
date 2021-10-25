package shared

type Evaluator func(scope *Scope, exp Expression) (Value, *Err)
