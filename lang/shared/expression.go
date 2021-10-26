package shared

const (
	ArgTypeValue    = "value"
	ArgTypeExp      = "expression"
	ArgTypeVariable = "variable"
	ArgTypeArgument = "argument"
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

func (v ArgValue) Type() string      { return ArgTypeValue }
func (v ArgVariable) Type() string   { return ArgTypeVariable }
func (v ArgExpression) Type() string { return ArgTypeExp }
func (v ArgArgument) Type() string   { return ArgTypeArgument }

type Expression struct {
	Operation string
	Arguments []ExpArgument
	Errors    []Err
}
