package shared

import "glisp/lang/errors"

type ArgType string

const (
	ArgTypeValue    ArgType = "value"
	ArgTypeExp      ArgType = "expression"
	ArgTypeVariable ArgType = "variable"
	ArgTypeArgument ArgType = "argument"
)

type ExpArgument interface {
	Type() ArgType
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

func (v ArgValue) Type() ArgType      { return ArgTypeValue }
func (v ArgVariable) Type() ArgType   { return ArgTypeVariable }
func (v ArgExpression) Type() ArgType { return ArgTypeExp }
func (v ArgArgument) Type() ArgType   { return ArgTypeArgument }

type Expression struct {
	Operation string
	Arguments []ExpArgument
	Errors    []errors.Err
}
