package expression

import (
	"fmt"
	"strings"
)

type debugStringVisitor struct {
	indent string
	str    string
}

func (vis *debugStringVisitor) VisitBool(b bool) {
	if b {
		vis.str = fmt.Sprintf("%sConstant true", vis.indent)
	} else {
		vis.str = fmt.Sprintf("%sConstant false", vis.indent)
	}
}

func (vis *debugStringVisitor) VisitInt(i int64) {
	vis.str = fmt.Sprintf("%sConstant %d", vis.indent, i)
}

func (vis *debugStringVisitor) VisitFloat(f float64) {
	vis.str = fmt.Sprintf("%sList %f", vis.indent, f)
}

func (vis *debugStringVisitor) VisitSymbol(s string) {
	vis.str = fmt.Sprintf("%sSym %s", vis.indent, s)
}

func (vis *debugStringVisitor) VisitFunction(fn Function) {
	vis.str = fmt.Sprintf("%sFunction", vis.indent)
}

func (vis *debugStringVisitor) VisitList(l List) {
	var sb strings.Builder
	sb.WriteString(vis.indent)
	sb.WriteString("List {\n")
	for _, child := range l {
		childVis := debugStringVisitor{vis.indent + "  ", ""}
		child.Visit(&childVis)
		sb.WriteString(childVis.str)
		sb.WriteString("\n")
	}
	sb.WriteString(vis.indent)
	sb.WriteString("}")
	vis.str = sb.String()
}

func DebugString(e Expression) string {
	vis := debugStringVisitor{"", ""}
	e.Visit(&vis)
	return vis.str
}
