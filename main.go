package main

import (
	"fmt"
	"glisp/lang"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		file, err := os.Open("samples/factorial.glisp")
		if err != nil {
			log.Fatal(err)
		}
		lang.Run(file)
		return
	}
	switch os.Args[1] {
	case "eval":
		if len(os.Args) < 3 {
			panic("Missing expression to eval")
		}
		exp := os.Args[2]
		fmt.Printf("glisp<< %s\n", exp)
		res := lang.EvalExp(strings.NewReader(exp))
		fmt.Println()
		fmt.Printf("glsip>> %v", res)
	case "run":
		if len(os.Args) < 3 {
			panic("Missing filepath")
		}
		file, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		lang.Run(file)
	default:
		panic("unknown command" + os.Args[1])
	}
}
