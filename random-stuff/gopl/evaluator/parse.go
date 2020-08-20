package evaluator

import (
	"fmt"
	"text/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

type lexPanic string

func (lex *lexer) next() {
	lex.token = lex.scan.Scan()
}

func (lex *lexer) text() string {
	return lex.scan.TokenText()
}

// Parse function parses the input string as an arithmetic expression
//
//	expr = num							a literal number eg., 1.732
//		|  id							a variable name eg., x
//		|  id '(' expr ',' ... ')'		a function call
//		|  '-' expr						a unary operator (+-)
//		|  expr '+' expr				a binary operator (+-*/)
//
func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// do not panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()
}
