package builtins

import (
	"errors"
	"fmt"
	"github.com/priyendra/dlisp/expression"
)

var Append BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		if len(args) != 2 {
			return nil, errors.New("Append operator requires exactly two args")
		}
		if expression.ToType(args[0]) != expression.LIST {
			return nil, errors.New("0th arg of append operator must be list")
		}
		return append(expression.AsList(args[0]), args[1]), nil
	},
}

var Car BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		if len(args) == 0 {
			return nil, errors.New("Car operator requires exactly one arg")
		}
		if expression.ToType(args[0]) != expression.LIST {
			return nil, errors.New("The arg of car operator must be list")
		}
		return expression.AsList(args[0])[0], nil
	},
}

var Cdr BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		if len(args) == 0 {
			return nil, errors.New("Cdr operator requires exactly one arg")
		}
		if expression.ToType(args[0]) != expression.LIST {
			return nil, errors.New("The arg of cdr operator must be list")
		}
		return expression.List(expression.AsList(args[0])[1:]), nil
	},
}

var Cons BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		if len(args) != 2 {
			return nil, errors.New("Cons operator requires exactly two args")
		}
		return expression.List(args), nil
	},
}

var Len BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		if len(args) == 0 {
			return nil, errors.New("Len operator requires exactly one arg")
		}
		if expression.ToType(args[0]) != expression.LIST {
			return nil, errors.New("The arg of len operator must be list")
		}
		return expression.Int(len(expression.AsList(args[0]))), nil
	},
}

var Map BuiltinFn = BuiltinFn{
	func(env *expression.Environment, args []expression.Expression) (
		expression.Expression, error) {
		if len(args) != 2 {
			return nil, errors.New("map operator requires exactly two args")
		}
		if expression.ToType(args[0]) != expression.FUNCTION {
			return nil, errors.New("0th arg of map operator must be function")
		}
		if expression.ToType(args[1]) != expression.LIST {
			return nil, errors.New("1th arg of map operator must be list")
		}
		mapped := make([]expression.Expression, len(expression.AsList(args[1])))
		for i, e := range expression.AsList(args[1]) {
			var err error
			mapped[i], err = expression.AsFunction(args[0]).Eval(
				env,
				[]expression.Expression{e})
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Error mapping %d-th arg", i))
			}
		}
		return expression.List(mapped), nil
	},
}
