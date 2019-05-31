package expression

type Type int

const (
	INT = iota
	FLOAT
	SYMBOL
	FUNCTION
	LIST
)

type Expression interface {
	Visit(vis Visitor)
}

type Visitor interface {
	VisitInt(i int64)
	VisitFloat(f float64)
	VisitSymbol(s string)
	VisitFunction(fn Function)
	VisitList(l List)
}

type Int int64
type Float float64
type Symbol string
type Function interface {
	Expression
	Eval(args []Expression) (Expression, error)
}
type List []Expression

func (i Int) Visit(vis Visitor)    { vis.VisitInt(int64(i)) }
func (f Float) Visit(vis Visitor)  { vis.VisitFloat(float64(f)) }
func (s Symbol) Visit(vis Visitor) { vis.VisitSymbol(string(s)) }
func (l List) Visit(vis Visitor)   { vis.VisitList(l) }

func AsInt(e Expression) int64         { return int64(e.(Int)) }
func AsFloat(e Expression) float64     { return float64(e.(Float)) }
func AsSymbol(e Expression) string     { return string(e.(Symbol)) }
func AsFunction(e Expression) Function { return e.(Function) }
func AsList(e Expression) List         { return e.(List) }
