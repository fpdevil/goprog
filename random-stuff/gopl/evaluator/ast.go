package evaluator

// A Var identifies a variable such as x, y, a etc.,
type Var string

// A literal is a numeric constant such as 1.414, 1.732 etc.,
type literal float64

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval represents the value of this Expr in the environment env
	Eval(env Env) float64
}

// A unary represents a unary operator such as +x, -x
type unary struct {
	op rune // any of +, -
	x  Expr
}

// A binary represents a binary operator expression such as
// x + y, x * a, x / y etc.,
type binary struct {
	op   rune // any of +, -, *, /
	x, y Expr
}

// A call represents a function call expression such as
// sin(x), cos(x), sqrt(x) etc.,
type call struct {
	fn   string // any of `sin`, `cos`, `sqrt`
	args []Expr
}
