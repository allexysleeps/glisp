package core

import "fmt"

func logOperation(name string, args []string, result primitiveValue) {
  fmt.Println("_____________________")
  fmt.Printf("called: %s with:\n%v\nresult: %v\n", name, args, result)
  fmt.Println("_____________________")
}

func operationDef(args []string)primitiveValue {
  varName := args[0]
  varValue := args[1]
  result := convertRaw2Primitive(varValue)
  lexicalScope[varName] = result
  logOperation("def", args, result)
  return result
}

func operationBool(args []string)primitiveValue {
  result := convert2Bool(convertRaw2Primitive(args[0]))
  logOperation("bool", args, result)
  return result
}