package builtins

import (
	"errors"
	"math"

	"github.com/priyendra/dlisp/expression"
)

func genericFloatOp(
	args []expression.Expression,
	op func(in float64) float64) (expression.Expression, error) {
	if len(args) != 1 {
		return nil, errors.New("Arithmetic operator requires exactly one arg")
	}
	switch expression.ToType(args[0]) {
	case expression.INT:
		return expression.Float(op(float64(expression.AsInt(args[0])))), nil
	case expression.FLOAT:
		return expression.Float(op(expression.AsFloat(args[0]))), nil
	case expression.BOOL:
		fallthrough
	case expression.SYMBOL:
		fallthrough
	case expression.FUNCTION:
		fallthrough
	case expression.LIST:
		return nil, errors.New("Non-numeric arg to math operator")
	}
	return nil, nil
}

var Abs BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		if len(args) != 1 {
			return nil, errors.New("Abs requires exactly one arg")
		}
		switch expression.ToType(args[0]) {
		case expression.INT:
			x := expression.AsInt(args[0])
			if x < 0 {
				return expression.Int(-x), nil
			}
			return expression.Int(x), nil
		case expression.FLOAT:
			return expression.Float(math.Abs(expression.AsFloat(args[0]))), nil
		case expression.BOOL:
			fallthrough
		case expression.SYMBOL:
			fallthrough
		case expression.FUNCTION:
			fallthrough
		case expression.LIST:
			return nil, errors.New("Abs requires arg of numeric type")
		}
		return nil, nil
	},
}

var Acos BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Acos(in) },
		)
	},
}

var Acosh BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Acosh(in) },
		)
	},
}

var Asin BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Asin(in) },
		)
	},
}

var Asinh BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Asinh(in) },
		)
	},
}

var Atan BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Atan(in) },
		)
	},
}

var Atanh BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Atanh(in) },
		)
	},
}

var Cbrt BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Cbrt(in) },
		)
	},
}

var Ceil BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Ceil(in) },
		)
	},
}

var Cos BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Cos(in) },
		)
	},
}

var Cosh BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Cosh(in) },
		)
	},
}

var Erf BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Erf(in) },
		)
	},
}

var Erfc BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Erfc(in) },
		)
	},
}

var Erfcinv BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Erfcinv(in) },
		)
	},
}

var Erfinv BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Erfinv(in) },
		)
	},
}

var Exp BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Exp(in) },
		)
	},
}

var Exp2 BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Exp2(in) },
		)
	},
}

var Floor BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Floor(in) },
		)
	},
}

var Gamma BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Gamma(in) },
		)
	},
}

var Log BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Log(in) },
		)
	},
}

var Log10 BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Log10(in) },
		)
	},
}

var Log2 BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Log2(in) },
		)
	},
}

var Round BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Round(in) },
		)
	},
}

var Sin BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Sin(in) },
		)
	},
}

var Sinh BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Sinh(in) },
		)
	},
}

var Sqrt BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Sqrt(in) },
		)
	},
}

var Tan BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Tan(in) },
		)
	},
}

var Tanh BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Tanh(in) },
		)
	},
}

var Trunc BuiltinFn = BuiltinFn{
	func(env expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		return genericFloatOp(
			args,
			func(in float64) float64 { return math.Trunc(in) },
		)
	},
}

// These Go math functions are currently not implemented. We can add them on an
// as needed basis.
// func Atan2(y, x float64) float64
// func Copysign(x, y float64) float64
// func Dim(x, y float64) float64
// func Yn(n int, x float64) float64
// func Float32bits(f float32) uint32
// func Float32frombits(b uint32) float32
// func Float64frombits(b uint64) float64
// func Float64bits(f float64) uint64
// func Frexp(f float64) (frac float64, exp int)
// func Hypot(p, q float64) float64
// func Ilogb(x float64) int
// func Inf(sign int) float64
// func IsInf(f float64, sign int) bool
// func IsNaN(f float64) (is bool)
// func Jn(n int, x float64) float64
// func Ldexp(frac float64, exp int) float64
// func Lgamma(x float64) (lgamma float64, sign int)
// func Modf(f float64) (int float64, frac float64)
// func NaN() float64
// func Nextafter32(x, y float32) (r float32)
// func Nextafter(x, y float64) (r float64)
// func Pow(x, y float64) float64
// func Signbit(x float64) bool
// func Sincos(x float64) (sin, cos float64)
// func J0(x float64) float64
// func J1(x float64) float64
// func Y0(x float64) float64
// func Y1(x float64) float64
// func Expm1(x float64) float64
// func Log1p(x float64) float64
// func Logb(x float64) float64
