package shared

type Evaluator func(exp Exp) interface{}

type Variable struct {
	Name  string
	Value string
}

func (v *Variable) Val() string {
	return v.Value
}
