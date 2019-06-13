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
		vis.err = errors.New("Could not eval second child of define expression: " +
			DebugString(l[2]))
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

func (vis *evalVisitor) handleLambda(l List) {
	if len(l) != 3 {
		vis.val = nil
		vis.err = errors.New("Lambda must have exactly two args")
		return
	}
	if ToType(l[1]) != LIST {
		vis.val = nil
		vis.err = errors.New("2nd arg to lambda must be a list of symbols")
		return
	}
	args := AsList(l[1])
	for _, a := range args {
		if ToType(a) != SYMBOL {
			vis.val = nil
			vis.err = errors.New("2nd arg to lambda must be a list of symbols")
			return
		}
	}
	vis.val = UserFunction{args, l[2]}
	vis.err = nil
}

// We need to support if as a special case to ensure that we only evaluate the
// branch of the if that matches. Recursive functions do not terminate otherwise
// at the base case if we perenially evaluate both branches.
func (vis *evalVisitor) handleIf(l List) {
	if len(l) != 4 {
		vis.val = nil
		vis.err = errors.New(
			"If expression must have the syntax (if condition then else)")
		return
	}
	condition, err := Eval(l[1], vis.env)
	if err != nil {
		vis.val = nil
		vis.err = errors.New("Could not eval if condition: " + DebugString(l[1]))
		return
	}
	if ToType(condition) != BOOL {
		vis.val = nil
		vis.err = errors.New("If condition must eval to boolean type")
		return
	}
	if AsBool(condition) {
		vis.val, vis.err = Eval(l[2], vis.env)
		return
	}
	vis.val, vis.err = Eval(l[3], vis.env)
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
	if ToType(l[0]) == SYMBOL && AsSymbol(l[0]) == "lambda" {
		vis.handleLambda(l)
		return
	}
	if ToType(l[0]) == SYMBOL && AsSymbol(l[0]) == "if" {
		vis.handleIf(l)
		return
	}
	childVis := evalVisitor{vis.env, nil, nil}
	l[0].Visit(&childVis)
	if childVis.err != nil {
		vis.val = nil
		vis.err = errors.New("Could not eval first child of list expression: " +
			DebugString(l[0]))
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
				"Could not eval %d-th arg of list expression (%s): "+DebugString(child),
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
