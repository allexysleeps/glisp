package shared

const (
	TypeValue    = "value"
	TypeExp      = "expression"
	TypeVariable = "variable"
	TypeArgument = "argument"
)

type ExpArgument interface {
	Type() string
}

type ArgValue struct {
	Value
}

type ArgExpression struct {
	Value *Expression
}

type ArgVariable struct {
	Value string
}

type ArgArgument struct {
	Value string
}

func (v ArgValue) Type() string      { return TypeValue }
func (v ArgVariable) Type() string   { return TypeVariable }
func (v ArgExpression) Type() string { return TypeExp }
func (v ArgArgument) Type() string   { return TypeArgument }

type Expression struct {
	Operation string
	Arguments []ExpArgument
	Errors    []Err
}
