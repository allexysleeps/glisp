package core

import (
  "fmt"
  "strconv"
)

func convertRaw2Primitive(rawValue string) primitiveValue {
  if rawValue == strNullValue {
    return primitiveValue{ null: true }
  }
  if rawValue == strTrueValue || rawValue == strFalseValue {
    return primitiveValue{ bool: rawValue == strTrueValue }
  }
  if numValue, err := strconv.ParseFloat(rawValue, 64); err == nil {
    return primitiveValue{ number: numValue }
  }
  if rawValue[0] == '"' && rawValue[len(rawValue) - 1] == '"' {
    return primitiveValue{ string: rawValue }
  }
  panic(fmt.Sprintf("Unexpected value: %v", rawValue))
}


func convert2Bool(value primitiveValue) primitiveValue {
  switch value.valueType {
  case typeNull: {
    return createBool(false)
  }
  case typeNumber: {
    return createBool(value.number != 0)
  }
  case typeString: {
    return createBool(value.string != "")
  }
  }
  return value // default if already boolean
}