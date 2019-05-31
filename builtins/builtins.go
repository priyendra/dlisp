package builtins

import (
	"errors"

	"github.com/priyendra/dlisp/expression"
)

type BuiltinFn struct {
	fn func(args []expression.Expression) (expression.Expression, error)
}

func (fn BuiltinFn) Visit(vis expression.Visitor) { vis.VisitFunction(fn) }
func (fn BuiltinFn) Eval(
	args []expression.Expression) (expression.Expression, error) {
	return fn.fn(args)
}

var Identity BuiltinFn = BuiltinFn{
	func(args []expression.Expression) (expression.Expression, error) {
		if len(args) != 1 {
			return nil, errors.New("Identity requires exactly one arg")
		}
		return args[0], nil
	},
}
