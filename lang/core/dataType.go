package core

type valueType int64

const (
  typeNull valueType = iota
  typeBool
  typeNumber
  typeString
)

type primitiveValue struct {
  bool bool
  number float64
  string string
  null bool
  valueType valueType
}

const (
  strTrueValue  = "true"
  strFalseValue = "false"
  strNullValue  = "null"
)

func createBool(value bool) primitiveValue {
  return primitiveValue{ bool: value, valueType: typeBool }
}