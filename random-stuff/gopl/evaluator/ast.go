// AST for Expression Evaluator

package evaluator

// An Expr us an arithmetic expression.
type Expr interface {
	// Eval returns the value if this Expr in the environment env
	Eval(env Env) float64
}

// A Var identifies a variable like x, y etc.,
type Var string

// A literal represents a numeric constant like 2.71728 etc.,
type literal float64

// A unary represents a unary operator expression like -x, +x etc.,
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// A binary represents a binary operator expression like x + y etc.,
type binary struct {
	op   rune // one if '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression like sin(x), cos(x) etc.,
type call struct {
	fn   string // one of 'pow', 'sin', 'sqrt'
	args []Expr
}
