package value

import "fmt"

type Type int

const (
	INT = iota
	FLOAT
	FUNCTION
)

type Value interface {
	Visit(vis Visitor)
}

type Visitor interface {
	VisitInt(i int64)
	VisitFloat(f float64)
	VisitFunction(fn Function)
}

// These are the three value types
type Int int64
type Float float64
type Function interface {
	Value
	Eval(args []Value) (Value, error)
}

func (i Int) Visit(vis Visitor)   { vis.VisitInt(int64(i)) }
func (f Float) Visit(vis Visitor) { vis.VisitFloat(float64(f)) }

type ToTypeVisitor struct {
	vtype Type
}

func (vis *ToTypeVisitor) VisitInt(i int64)          { vis.vtype = INT }
func (vis *ToTypeVisitor) VisitFloat(f float64)      { vis.vtype = FLOAT }
func (vis *ToTypeVisitor) VisitFunction(fn Function) { vis.vtype = FUNCTION }

type ToStringVisitor struct {
	str string
}

func (vis *ToStringVisitor) VisitInt(i int64) {
	vis.str = fmt.Sprintf("%d", i)
}
func (vis *ToStringVisitor) VisitFloat(f float64) {
	vis.str = fmt.Sprintf("%f", f)
}
func (vis *ToStringVisitor) VisitFunction(fn Function) {
	vis.str = "Function"
}

func ToType(v Value) Type {
	vis := ToTypeVisitor{}
	v.Visit(&vis)
	return vis.vtype
}

func ToString(v Value) string {
	vis := ToStringVisitor{}
	v.Visit(&vis)
	return vis.str
}
