package parse

import (
	"fmt"
	"glisp/lang/expression"
	"glisp/lang/parse/symbols"
	"io"
	"text/scanner"
)

type lexer struct {
	scan scanner.Scanner
	token rune
}

type lexPanic string

func (lex *lexer) next() { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func Parse(input io.Reader) []expression.Exp {
	lex := new(lexer)
	var expressions []expression.Exp
	lex.scan.Init(input)
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats | scanner.ScanStrings
	lex.next()
out:
	for {
		switch lex.token {
		case symbols.ParOpen:
			expressions = append(expressions, parseExp(lex))
		case scanner.EOF:
			break out
		default:
			panic(fmt.Errorf("unexpected symbol: %s", lex.text()))
		}
	}
	return expressions
}

func parseExp(lex *lexer) expression.Exp {
	var exp expression.Exp

	parser(lex, &exp)
	return exp
}

func parser(lex *lexer, exp *expression.Exp) *expression.Exp {
	lex.next()

	if lex.token == symbols.ParClose {
		lex.next()
		return nil
	}
	exp.Operation = parseOperation(lex)

	var args []expression.ExpArgument
out:
	for {
		switch lex.token {
		case symbols.ParOpen:
			args = append(args, expression.ArgExp{ Val: parser(lex, exp) })
		case symbols.ParClose:
			lex.next()
			break out
		default:
			args = append(args, expression.ArgValue{ Val: lex.text() })
			lex.next()
		}
	}
	return exp
}

func parseOperation(lex *lexer) string {
	op := lex.text()
	lex.next()
	return op
}