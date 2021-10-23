package main

import (
	"fmt"
	"glisp/lang"
	"log"
	"os"
	"strings"
)

func main() {
	switch os.Args[1] {
	case "eval":
		if len(os.Args) < 3 {
			panic("Missing expression to eval")
		}
		exp := os.Args[2]
		fmt.Printf("glisp<< %s\n", exp)
		lang.EvalExp(strings.NewReader(exp))
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
		file, err := os.Open("samples/basics.glisp")
		if err != nil {
			log.Fatal(err)
		}
		lang.Run(file)
	}
}
