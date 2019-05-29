package builtins

import (
	"errors"
	"github.com/priyendra/dlisp/value"
)

type BuiltinFn struct {
	fn func(args []value.Value) (value.Value, error)
}

func (fn BuiltinFn) Visit(vis value.Visitor) { vis.VisitFunction(fn) }
func (fn BuiltinFn) Eval(args []value.Value) (value.Value, error) {
	return fn.fn(args)
}

var Identity BuiltinFn = BuiltinFn{
	func(args []value.Value) (value.Value, error) {
		if len(args) != 1 {
			return nil, errors.New("Identity requires exactly one argument")
		}
		return args[0], nil
	},
}
