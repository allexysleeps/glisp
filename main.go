package main

import (
	"glisp/lang/core"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("samples/basics.glisp")
	if err != nil {
		log.Fatal(err)
	}
	core.Compile(b)
}