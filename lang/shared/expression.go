package shared

const (
	TypeValue    = "value"
	TypeExp      = "expression"
	TypeVariable = "variable"
)

type ExpArgument interface {
	Type() string
}

type ArgValue struct {
	Value
}

type ArgExp struct {
	Value *Exp
}

type ArgVariable struct {
	Value string
}

func (v ArgValue) Type() string    { return TypeValue }
func (v ArgVariable) Type() string { return TypeVariable }
func (v ArgExp) Type() string      { return TypeExp }

type Exp struct {
	Operation string
	Arguments []ExpArgument
}
