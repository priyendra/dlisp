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

var If BuiltinFn = BuiltinFn{
	func(args []expression.Expression) (expression.Expression, error) {
		if len(args) != 3 {
			return nil, errors.New("If requires exactly three args")
		}
		if expression.ToType(args[0]) != expression.BOOL {
			return nil, errors.New("If: 0th arg must be of bool type")
		}
		if expression.AsBool(args[0]) {
			return args[1], nil
		}
		return args[2], nil
	},
}
