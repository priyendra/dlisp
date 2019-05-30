package expression

import (
	"github.com/priyendra/dlisp/builtins"
	"github.com/priyendra/dlisp/value"
	"math"
)

type Environment struct {
	names map[string]value.Value
}

func StdEnv() Environment {
	env := Environment{
		map[string]value.Value{},
	}
	env.names["identity"] = builtins.Identity
	env.names["+"] = builtins.Plus
	env.names["-"] = builtins.Minus
	env.names["*"] = builtins.Multiply
	env.names["/"] = builtins.Divide
	env.names["%"] = builtins.Mod
	env.names["pi"] = value.Float(math.Pi)
	return env
}
