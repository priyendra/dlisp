package builtins

import (
	"errors"
	"fmt"
	"github.com/priyendra/dlisp/value"
)

type plusVisitor struct {
	operandIndex int
	intSum       int64
	floatSum     float64
	allInts      bool
	err          error
}

func (vis *plusVisitor) VisitInt(i int64) {
	if vis.allInts {
		vis.intSum += i
	} else {
		vis.floatSum += float64(i)
	}
	vis.operandIndex++
}

func (vis *plusVisitor) VisitFloat(f float64) {
	if vis.allInts {
		vis.allInts = false
		vis.floatSum = float64(vis.intSum)
	}
	vis.floatSum += f
	vis.operandIndex++
}

func newPlusVisitor() plusVisitor {
	return plusVisitor{0, int64(0), float64(0.0), true, nil}
}

func (vis *plusVisitor) VisitFunction(fn value.Function) {
	vis.err = errors.New(fmt.Sprintf(
		"Plus expected numeric type, found function at operand %d",
		vis.operandIndex))
	vis.operandIndex++
}

var Plus BuiltinFn = BuiltinFn{
	func(args []value.Value) (value.Value, error) {
		if len(args) < 1 {
			return nil, errors.New("Plus requires at least one argument")
		}
		plusVis := newPlusVisitor()
		for _, a := range args {
			a.Visit(&plusVis)
			if plusVis.err != nil {
				return nil, plusVis.err
			}
		}
		if plusVis.allInts {
			return value.Int(plusVis.intSum), nil
		}
		return value.Float(plusVis.floatSum), nil
	},
}
