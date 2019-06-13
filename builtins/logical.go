package builtins

import (
	"errors"

	"github.com/priyendra/dlisp/expression"
)

var LogicalAnd BuiltinFn = BuiltinFn{
	func(args []expression.Expression) (
		expression.Expression, error) {
		if len(args) != 2 {
			return nil, errors.New("Logical operator AND requires exactly two args")
		}
		if expression.ToType(args[0]) != expression.BOOL ||
			expression.ToType(args[1]) != expression.BOOL {
			return nil, errors.New(
				"Logical operator AND requires both args to have boolean type")
		}
		return expression.Bool(
			expression.AsBool(args[0]) && expression.AsBool(args[1])), nil
	},
}

var LogicalOr BuiltinFn = BuiltinFn{
	func(args []expression.Expression) (
		expression.Expression, error) {
		if len(args) != 2 {
			return nil, errors.New("Logical operator AND requires exactly two args")
		}
		if expression.ToType(args[0]) != expression.BOOL ||
			expression.ToType(args[1]) != expression.BOOL {
			return nil, errors.New(
				"Logical operator OR requires both args to have boolean type")
		}
		return expression.Bool(
			expression.AsBool(args[0]) || expression.AsBool(args[1])), nil
	},
}

var LogicalNot BuiltinFn = BuiltinFn{
	func(args []expression.Expression) (
		expression.Expression, error) {
		if len(args) != 1 {
			return nil, errors.New("Logical operator NOT requires exactly one arg")
		}
		if expression.ToType(args[0]) != expression.BOOL {
			return nil, errors.New(
				"Logical operator NOT requires argg to have boolean type")
		}
		return expression.Bool(!expression.AsBool(args[0])), nil
	},
}
