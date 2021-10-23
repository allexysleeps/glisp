package main

import (
  "glisp/lang"
  "log"
  "os"
)

func main() {
  file, err := os.Open("samples/basics.glisp")
  if err != nil {
     log.Fatal(err)
  }
  lang.Run(file)
}
