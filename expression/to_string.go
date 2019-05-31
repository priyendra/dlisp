package expression

import (
	"fmt"
	"strings"
)

type stringifyVisitor struct {
	str string
}

func (vis *stringifyVisitor) VisitBool(b bool) {
	if b {
		vis.str = "true"
	} else {
		vis.str = "false"
	}
}
func (vis *stringifyVisitor) VisitInt(i int64) {
	vis.str = fmt.Sprintf("%d", i)
}
func (vis *stringifyVisitor) VisitFloat(f float64) {
	vis.str = fmt.Sprintf("%f", f)
}
func (vis *stringifyVisitor) VisitSymbol(s string) {
	vis.str = fmt.Sprintf("Symbol[%s]", s)
}
func (vis *stringifyVisitor) VisitFunction(fn Function) {
	vis.str = "Function"
}
func (vis *stringifyVisitor) VisitList(l List) {
	var sb strings.Builder
	sb.WriteString("(")
	for i, a := range l {
		childVis := stringifyVisitor{}
		a.Visit(&childVis)
		sb.WriteString(childVis.str)
		if i < len(l)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	vis.str = sb.String()
}

func ToString(e Expression) string {
	vis := stringifyVisitor{}
	e.Visit(&vis)
	return vis.str
}
