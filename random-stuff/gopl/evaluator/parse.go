package evaluator

import (
	"fmt"

	"github.com/hashicorp/hcl/hcl/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan()
}

func (lex *lexer) text() string {
	return lex.scan.TokenText()
}

type lexPanic string

//!+Parse

// Parse function parses the input string as an arithmetic expression
//
// expr = num						a literal number e.g., 2.71728
//		| id						a variable name e.g., x
//		| id '(' expr ',' ... ')'	a function call
//		| '-' expr					a unary operator (+-)
//		| expr '+' expr				a binary operator (+-*/)
//
func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic
			panic(x)
		}
	}()

}

//!-Parse
