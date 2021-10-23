package shared

import "strconv"

type Evaluator func(scope *Scope, exp Exp) interface{}

type Variable struct {
	Name  string
	Value string
}

func (v *Variable) Val() string {
	return v.Value
}

func (v *Variable) NumVal() (float64, error) {
	val, err := strconv.ParseFloat(v.Value, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}
