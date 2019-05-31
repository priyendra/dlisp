package expression

import (
	"errors"
	"fmt"
)

type evalVisitor struct {
	env *Environment
	val Expression
	err error
}

func (vis *evalVisitor) VisitInt(i int64) {
	vis.val = Int(i)
	vis.err = nil
}

func (vis *evalVisitor) VisitFloat(f float64) {
	vis.val = Float(f)
	vis.err = nil
}

func (vis *evalVisitor) VisitSymbol(s string) {
	val, found := vis.env.Names[s]
	if !found {
		vis.val = nil
		vis.err = errors.New(fmt.Sprintf("Undefined symbol: '%s'", s))
		return
	}
	vis.val = val
	vis.err = nil
}

func (vis *evalVisitor) VisitFunction(fn Function) {
	vis.val = fn
	vis.err = nil
}

func (vis *evalVisitor) VisitList(l List) {
	if len(l) < 1 {
		vis.err = errors.New("Empty list expression not supported")
		return
	}
	childVis := evalVisitor{vis.env, nil, nil}
	l[0].Visit(&childVis)
	if childVis.err != nil {
		vis.val = nil
		vis.err = errors.New("Could not eval first child of list expression")
		return
	}
	if ToType(childVis.val) != FUNCTION {
		vis.val = nil
		vis.err = errors.New(
			"First child of list function must eval to type function")
		return
	}
	args := make([]Expression, len(l)-1)
	for i, child := range l[1:] {
		childVis2 := evalVisitor{vis.env, nil, nil}
		child.Visit(&childVis2)
		if childVis2.err != nil {
			vis.err = errors.New(fmt.Sprintf(
				"Could not eval %d-th arg of list expression (%s)",
				i+1, childVis2.err.Error()))
			return
		}
		args[i] = childVis2.val
	}
	vis.val, vis.err = AsFunction(childVis.val).Eval(args)
}

func Eval(e Expression, env *Environment) (Expression, error) {
	vis := evalVisitor{env, nil, nil}
	e.Visit(&vis)
	return vis.val, vis.err
}
