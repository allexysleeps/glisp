package core

var operationsMap = map[string]func(args []string)primitiveValue {
  "def": operationDef,
  "bool": operationBool,
}