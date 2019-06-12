package expression

import (
	"errors"
	"fmt"
)

type evalVisitor struct {
	env Environment
	val Expression
	err error
}

func (vis *evalVisitor) VisitBool(b bool) {
	vis.val = Bool(b)
	vis.err = nil
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
	vis.val, vis.err = vis.env.LookupVar(s)
}

func (vis *evalVisitor) VisitFunction(fn Function) {
	vis.val = fn
	vis.err = nil
}

func (vis *evalVisitor) handleDefine(l List) {
	if len(l) != 3 {
		vis.val = nil
		vis.err = errors.New("Define must have exactly two args")
		return
	}
	if ToType(l[1]) != SYMBOL {
		vis.val = nil
		vis.err = errors.New("First arg to define must be symbol")
		return
	}
	childVis := evalVisitor{vis.env, nil, nil}
	l[2].Visit(&childVis)
	if childVis.err != nil {
		vis.val = nil
		vis.err = errors.New("Could not eval second child of define expression")
		return
	}
	vis.env.SetVar(AsSymbol(l[1]), childVis.val)
	vis.val = childVis.val
	vis.err = nil
	return
}

func (vis *evalVisitor) handleQuote(l List) {
	if len(l) != 2 {
		vis.val = nil
		vis.err = errors.New("Quote must have exactly one arg")
		return
	}
	vis.val = l[1]
	vis.err = nil
}

func (vis *evalVisitor) VisitList(l List) {
	if len(l) < 1 {
		vis.err = errors.New("Empty list expression not supported")
		return
	}
	if ToType(l[0]) == SYMBOL && AsSymbol(l[0]) == "define" {
		vis.handleDefine(l)
		return
	}
	if ToType(l[0]) == SYMBOL && AsSymbol(l[0]) == "quote" {
		vis.handleQuote(l)
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
	vis.val, vis.err = AsFunction(childVis.val).Eval(vis.env, args)
}

func Eval(e Expression, env Environment) (Expression, error) {
	vis := evalVisitor{env, nil, nil}
	e.Visit(&vis)
	return vis.val, vis.err
}
