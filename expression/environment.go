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
	env.names["pi"] = value.Float(math.Pi)
	return env
}
