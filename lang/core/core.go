package core

import (
  "fmt"
  "strings"
)

const (
  parenthesisOpen = '('
  parenthesisClose = ')'
)

var expressions [][]byte

func Compile(srcCode []byte) {
  var start int
  for i := 0; i < len(srcCode); i++ {
    if srcCode[i] == parenthesisOpen {
      start = i
      continue
    }
    if srcCode[i] == parenthesisClose {
      expressions = append(expressions, srcCode[start + 1:i])
      start = 0
    }
  }

  for _, exp := range expressions {
    args := strings.Fields(string(exp))
    Keywords[args[0]](args[0:])
  }
  logLexicalEnv()
}

func logLexicalEnv() {
  for key, value := range LexicalScope {
    fmt.Printf("%s: '%s'\n", key, value)
  }
}