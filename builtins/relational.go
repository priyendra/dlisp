package builtins

import (
	"errors"

	"github.com/priyendra/dlisp/expression"
)

func genericRelationalOp(
	args []expression.Expression,
	intFn func(a int64, b int64) bool,
	floatFn func(a float64, b float64) bool) (expression.Expression, error) {
	if len(args) != 2 {
		return nil, errors.New("Relational operator requires exactly two args")
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
		return nil, errors.New("Non-numeric arg to relational operator")
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
		return nil, errors.New("Non-numeric arg to relational operator")
	}
	if allInts {
		return expression.Bool(intFn(aInt, bInt)), nil
	}
	return expression.Bool(floatFn(aFloat, bFloat)), nil
}

var GreaterThan BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericRelationalOp(
			args,
			func(a int64, b int64) bool { return a > b },
			func(a float64, b float64) bool { return a > b },
		)
	},
}

var GreaterEqual BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericRelationalOp(
			args,
			func(a int64, b int64) bool { return a >= b },
			func(a float64, b float64) bool { return a >= b },
		)
	},
}

var LessThan BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericRelationalOp(
			args,
			func(a int64, b int64) bool { return a < b },
			func(a float64, b float64) bool { return a < b },
		)
	},
}

var LessEqual BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericRelationalOp(
			args,
			func(a int64, b int64) bool { return a <= b },
			func(a float64, b float64) bool { return a <= b },
		)
	},
}

var Equal BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		allNumeric := true
		allBool := true
		for _, a := range args {
			if expression.ToType(a) != expression.INT &&
				expression.ToType(a) != expression.FLOAT {
				allNumeric = false
			}
			if expression.ToType(a) != expression.BOOL {
				allBool = false
			}
		}
		if allNumeric {
			return genericRelationalOp(
				args,
				func(a int64, b int64) bool { return a == b },
				func(a float64, b float64) bool { return a == b },
			)
		}
		if allBool {
			return expression.Bool(
					expression.AsBool(args[0]) == expression.AsBool(args[1])),
				nil
		}
		return nil, errors.New("Unexpected type in Equals")
	},
}

var NotEqual BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		allNumeric := true
		allBool := true
		for _, a := range args {
			if expression.ToType(a) != expression.INT &&
				expression.ToType(a) != expression.FLOAT {
				allNumeric = false
			}
			if expression.ToType(a) != expression.BOOL {
				allBool = false
			}
		}
		if allNumeric {
			return genericRelationalOp(
				args,
				func(a int64, b int64) bool { return a != b },
				func(a float64, b float64) bool { return a != b },
			)
		}
		if allBool {
			return expression.Bool(
				expression.AsBool(args[0]) != expression.AsBool(args[1])), nil
		}
		return nil, errors.New("Unexpected type in NotEquals")
	},
}
