package parse

import (
	"fmt"
	"glisp/lang/parse/symbols"
	"glisp/lang/shared"
	"io"
	"text/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func Parse(input io.Reader) []shared.Exp {
	lex := new(lexer)
	var expressions []shared.Exp
	lex.scan.Init(input)
	lex.scan.Mode =
		scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats | scanner.ScanStrings | scanner.SkipComments
	lex.next()
out:
	for {
		switch lex.token {
		case symbols.ParOpen:
			expressions = append(expressions, *parseExp(lex))
		case scanner.EOF:
			break out
		default:
			panic(fmt.Errorf("unexpected symbol: %s", lex.text()))
		}
	}
	return expressions
}

func parseExp(lex *lexer) *shared.Exp {
	var exp shared.Exp

	lex.next()

	if lex.token == symbols.ParClose {
		lex.next()
		return nil
	}
	exp.Operation = parseOperation(lex)

	var args []shared.ExpArgument
out:
	for {
		fmt.Println(lex.text())
		switch lex.token {
		case symbols.ParOpen:
			args = append(args, shared.ArgExp{Value: parseExp(lex)})
		case symbols.ParClose:
			lex.next()
			break out
		default:
			args = append(args, shared.ArgValue{Value: lex.text()})
			lex.next()
		}
	}
	exp.Arguments = args
	return &exp
}

func parseOperation(lex *lexer) string {
	op := lex.text()
	lex.next()
	return op
}
