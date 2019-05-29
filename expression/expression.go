package expression

import (
	"errors"
	"fmt"
	"github.com/priyendra/dlisp/value"
)

type Type int

const (
	INT_LITERAL = iota
	FLOAT_LITERAL
	SYMBOL
	COMPOUND
)

type Expression interface {
	Visit(vis Visitor)
}

type Visitor interface {
	VisitIntLiteral(i int64)
	VisitFloatLiteral(f float64)
	VisitSymbol(s string)
	VisitCompound(c Compound)
}

type IntLiteral int64
type FloatLiteral float64
type Symbol string
type Compound []Expression

func (i IntLiteral) Visit(vis Visitor)   { vis.VisitIntLiteral(int64(i)) }
func (f FloatLiteral) Visit(vis Visitor) { vis.VisitFloatLiteral(float64(f)) }
func (s Symbol) Visit(vis Visitor)       { vis.VisitSymbol(string(s)) }
func (c Compound) Visit(vis Visitor)     { vis.VisitCompound(c) }

type EvalVisitor struct {
	env Environment
	val value.Value
	err error
}

func (vis *EvalVisitor) VisitIntLiteral(i int64) {
	vis.val = value.Int(i)
	vis.err = nil
}

func (vis *EvalVisitor) VisitFloatLiteral(f float64) {
	vis.val = value.Float(f)
	vis.err = nil
}

func (vis *EvalVisitor) VisitSymbol(s string) {
	val, found := vis.env.names[s]
	if !found {
		vis.val = nil
		vis.err = errors.New(fmt.Sprintf("Undefined symbol: '%s'", s))
		return
	}
	vis.val = val
	vis.err = nil
}

func (vis *EvalVisitor) VisitCompound(c Compound) {
	if len(c) < 1 {
		vis.err = errors.New("Empty compound expression not supported")
		return
	}
	childVis := EvalVisitor{vis.env, nil, nil}
	c[0].Visit(&childVis)
	if childVis.err != nil {
		vis.val = nil
		vis.err = errors.New("Could not eval first child of list expression")
		return
	}
	if value.ToType(childVis.val) != value.FUNCTION {
		vis.val = nil
		vis.err = errors.New(
			"First child of list function must eval to type function")
		return
	}
	args := make([]value.Value, len(c)-1)
	for i, child := range c[1:] {
		childVis2 := EvalVisitor{vis.env, nil, nil}
		child.Visit(&childVis2)
		if childVis2.err != nil {
			vis.err = errors.New(
				fmt.Sprintf("Could not eval %d-th arg of list expression", i-1))
			return
		}
		args[i] = childVis2.val
	}
	vis.val, vis.err = childVis.val.(value.Function).Eval(args)
}

func Eval(e Expression) (value.Value, error) {
	vis := EvalVisitor{StdEnv(), nil, nil}
	e.Visit(&vis)
	return vis.val, vis.err
}
