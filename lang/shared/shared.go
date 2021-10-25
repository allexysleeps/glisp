package shared

type Evaluator func(scope *Scope, exp Exp) Value

type Variable struct {
	Name string
	Value
}
