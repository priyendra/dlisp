package builtins

import (
	"errors"
	"math"

	"github.com/priyendra/dlisp/expression"
)

func genericArithemeticOp(
	args []expression.Expression,
	intFn func(a int64, b int64) int64,
	floatFn func(a float64, b float64) float64) (expression.Expression, error) {
	if len(args) != 2 {
		return nil, errors.New("Arithmetic operator requires exactly two args")
	}
	allInts := true
	var aInt, bInt int64
	var aFloat, bFloat float64
	switch expression.ToType(args[0]) {
	case expression.INT:
		aInt = expression.AsInt(args[0])
		aFloat = float64(aInt)
	case expression.FLOAT:
		allInts = false
		aFloat = expression.AsFloat(args[0])
	case expression.BOOL:
		fallthrough
	case expression.SYMBOL:
		fallthrough
	case expression.FUNCTION:
		fallthrough
	case expression.LIST:
		return nil, errors.New("Non-numeric arg to arithmetic operator")
	}
	switch expression.ToType(args[1]) {
	case expression.INT:
		bInt = expression.AsInt(args[1])
		bFloat = float64(bInt)
	case expression.FLOAT:
		allInts = false
		bFloat = expression.AsFloat(args[1])
	case expression.BOOL:
		fallthrough
	case expression.SYMBOL:
		fallthrough
	case expression.FUNCTION:
		fallthrough
	case expression.LIST:
		return nil, errors.New("Non-numeric arg to arithmetic operator")
	}
	if allInts {
		return expression.Int(intFn(aInt, bInt)), nil
	}
	return expression.Float(floatFn(aFloat, bFloat)), nil
}

var Plus BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericArithemeticOp(
			args,
			func(a int64, b int64) int64 { return a + b },
			func(a float64, b float64) float64 { return a + b },
		)
	},
}

var Minus BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericArithemeticOp(
			args,
			func(a int64, b int64) int64 { return a - b },
			func(a float64, b float64) float64 { return a - b },
		)
	},
}

var Multiply BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericArithemeticOp(
			args,
			func(a int64, b int64) int64 { return a * b },
			func(a float64, b float64) float64 { return a * b },
		)
	},
}

var Divide BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericArithemeticOp(
			args,
			func(a int64, b int64) int64 { return a / b },
			func(a float64, b float64) float64 { return a / b },
		)
	},
}

var Max BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericArithemeticOp(
			args,
			func(a int64, b int64) int64 {
				if a > b {
					return a
				}
				return b
			},
			func(a float64, b float64) float64 { return math.Max(a, b) },
		)
	},
}

var Min BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericArithemeticOp(
			args,
			func(a int64, b int64) int64 {
				if a < b {
					return a
				}
				return b
			},
			func(a float64, b float64) float64 { return math.Min(a, b) },
		)
	},
}

var Mod BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericArithemeticOp(
			args,
			func(a int64, b int64) int64 { return a % b },
			func(a float64, b float64) float64 { return math.Mod(a, b) },
		)
	},
}
