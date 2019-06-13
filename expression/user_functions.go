package expression

import (
	"errors"
	"fmt"
)

type UserFunction struct {
	args List
	body Expression
	env  Environment
}

type NestedEnvironment struct {
	base Environment
	vars map[string]Expression
}

func (env NestedEnvironment) LookupVar(name string) (Expression, error) {
	val, ok := env.vars[name]
	if !ok {
		return env.base.LookupVar(name)
	}
	return val, nil
}

func (env *NestedEnvironment) SetVar(name string, value Expression) {
	env.vars[name] = value
}

func (fn UserFunction) Visit(vis Visitor) { vis.VisitFunction(fn) }

func (fn UserFunction) Eval(args []Expression) (
	Expression, error) {
	if len(args) != len(fn.args) {
		return nil, errors.New(fmt.Sprintf(
			"UserFunction expected %d args, found %d", len(fn.args), len(args)))
	}
	newEnv := NestedEnvironment{fn.env, map[string]Expression{}}
	for i, a := range fn.args {
		newEnv.vars[AsSymbol(a)] = args[i]
	}
	return Eval(fn.body, &newEnv)
}
