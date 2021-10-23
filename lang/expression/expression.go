package expression

const (
	TypeValue = "value"
	TypeExp = "expression"
)

type ExpArgument interface {
	Type() string
}

type ArgValue struct {
	Val string
}

type ArgExp struct {
	Val *Exp
}

func (v ArgValue) Type() string { return TypeValue }
func (v ArgExp) Type() string { return TypeExp }

type Exp struct {
	Operation string
	Arguments []ExpArgument
}