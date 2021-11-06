package main

import (
	"glisp/lang"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		file, err := os.Open("samples/sumList.glisp")
		if err != nil {
			log.Fatal(err)
		}
		lang.Run(file)
		return
	}

	requiredArgsLen := 3

	switch os.Args[1] {
	case "eval":
		if len(os.Args) < requiredArgsLen {
			panic("Missing expression to eval")
		}
		exp := os.Args[2]
		log.Printf("glisp<< %s\n", exp)
		res := lang.EvalExp(strings.NewReader(exp))
		log.Println()
		log.Printf("glsip>> %v", res)
	case "run":
		if len(os.Args) < requiredArgsLen {
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
