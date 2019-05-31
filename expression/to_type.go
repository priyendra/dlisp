package expression

// Convert an expression to its type
type typeVisitor struct {
	etype Type
}

func (vis *typeVisitor) VisitInt(i int64)          { vis.etype = INT }
func (vis *typeVisitor) VisitFloat(f float64)      { vis.etype = FLOAT }
func (vis *typeVisitor) VisitSymbol(s string)      { vis.etype = SYMBOL }
func (vis *typeVisitor) VisitFunction(fn Function) { vis.etype = FUNCTION }
func (vis *typeVisitor) VisitList(l List)          { vis.etype = LIST }

func ToType(e Expression) Type {
	vis := typeVisitor{}
	e.Visit(&vis)
	return vis.etype
}
