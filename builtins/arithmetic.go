package builtins

import (
	"errors"
	"github.com/priyendra/dlisp/value"
	"math"
)

func genericArithemeticOperator(
	args []value.Value,
	intFn func(a int64, b int64) int64,
	floatFn func(a float64, b float64) float64) (value.Value, error) {
	if len(args) != 2 {
		return nil, errors.New("Arithmetic operator requires exactly two arguments")
	}
	allInts := true
	var aInt, bInt int64
	var aFloat, bFloat float64
	switch value.ToType(args[0]) {
	case value.INT:
		aInt = int64(args[0].(value.Int))
		aFloat = float64(aInt)
	case value.FLOAT:
		allInts = false
		aFloat = float64(args[0].(value.Float))
	case value.FUNCTION:
		return nil, errors.New("Non-numeric argument to arithmetic operator")
	}
	switch value.ToType(args[1]) {
	case value.INT:
		bInt = int64(args[1].(value.Int))
		bFloat = float64(bInt)
	case value.FLOAT:
		allInts = false
		bFloat = float64(args[1].(value.Float))
	case value.FUNCTION:
		return nil, errors.New("Non-numeric argument to arithmetic operator")
	}
	if allInts {
		return value.Int(intFn(aInt, bInt)), nil
	}
	return value.Float(floatFn(aFloat, bFloat)), nil
}

var Plus BuiltinFn = BuiltinFn{
	func(args []value.Value) (value.Value, error) {
		return genericArithemeticOperator(
			args,
			func(a int64, b int64) int64 { return a + b },
			func(a float64, b float64) float64 { return a + b },
		)
	},
}

var Minus BuiltinFn = BuiltinFn{
	func(args []value.Value) (value.Value, error) {
		return genericArithemeticOperator(
			args,
			func(a int64, b int64) int64 { return a - b },
			func(a float64, b float64) float64 { return a - b },
		)
	},
}

var Multiply BuiltinFn = BuiltinFn{
	func(args []value.Value) (value.Value, error) {
		return genericArithemeticOperator(
			args,
			func(a int64, b int64) int64 { return a * b },
			func(a float64, b float64) float64 { return a * b },
		)
	},
}

var Divide BuiltinFn = BuiltinFn{
	func(args []value.Value) (value.Value, error) {
		return genericArithemeticOperator(
			args,
			func(a int64, b int64) int64 { return a / b },
			func(a float64, b float64) float64 { return a / b },
		)
	},
}

var Mod BuiltinFn = BuiltinFn{
	func(args []value.Value) (value.Value, error) {
		return genericArithemeticOperator(
			args,
			func(a int64, b int64) int64 { return a % b },
			func(a float64, b float64) float64 { return math.Mod(a, b) },
		)
	},
}
