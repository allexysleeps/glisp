package parse

import (
	"fmt"
	"io"
	"text/scanner"

	"glisp/lang/shared"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func Parse(input io.Reader) []shared.Expression {
	lex := new(lexer)
	var expressions []shared.Expression
	lex.scan.Init(input)
	lex.scan.Mode =
		scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats | scanner.ScanStrings | scanner.SkipComments
	lex.next()
out:
	for {
		switch lex.token {
		case symbolParOpen:
			expressions = append(expressions, *parseExp(lex))
		case scanner.EOF:
			break out
		default:
			panic(fmt.Errorf("unexpected symbol: %s", lex.text()))
		}
	}
	return expressions
}

func parseExp(lex *lexer) *shared.Expression {
	var exp shared.Expression

	lex.next()

	if lex.token == symbolParClose {
		lex.next()
		return nil
	}

	exp.Operation = parseOperation(lex)

	var args []shared.ExpArgument

	prefixes := make(map[string]bool)
out:
	for {
		switch lex.token {
		case symbolParOpen:
			args = append(args, shared.ArgExpression{Value: parseExp(lex)})
		case symbolParClose:
			lex.next()
			break out
		case symbolMinus:
			prefixes["minus"] = true
			lex.next()
		case symbolSquareBracketOpen:
			prefixes["squareOpen"] = true
			lex.next()
		case symbolSquareBracketClose:
			prefixes["squareOpen"] = false
			lex.next()
		default:
			val, ok := shared.CreateValue(lex.text())
			if !ok {
				if prefixes["squareOpen"] {
					args = append(args, shared.ArgArgument{Value: lex.text()})
				} else {
					args = append(args, shared.ArgVariable{Value: lex.text()})
				}
			} else {
				if val.Type() == shared.TypeNum && prefixes["minus"] {
					val, _ = shared.CreateValue("-" + val.StrVal())
					prefixes["minus"] = false
				}
				args = append(args, shared.ArgValue{Value: val})
			}
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
