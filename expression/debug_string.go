package expression

import (
	"fmt"
	"strings"
)

type DebugStringVisitor struct {
	indent string
	str    string
}

func (vis *DebugStringVisitor) VisitIntLiteral(i int64) {
	vis.str = fmt.Sprintf("%sLit %d", vis.indent, i)
}

func (vis *DebugStringVisitor) VisitFloatLiteral(f float64) {
	vis.str = fmt.Sprintf("%sLit %f", vis.indent, f)
}

func (vis *DebugStringVisitor) VisitSymbol(s string) {
	vis.str = fmt.Sprintf("%sSym %s", vis.indent, s)
}

func (vis *DebugStringVisitor) VisitCompound(c Compound) {
	var sb strings.Builder
	sb.WriteString(vis.indent)
	sb.WriteString("Compound {\n")
	for _, child := range c {
		childVis := DebugStringVisitor{vis.indent + "  ", ""}
		child.Visit(&childVis)
		sb.WriteString(childVis.str)
		sb.WriteString("\n")
	}
	sb.WriteString(vis.indent)
	sb.WriteString("}")
	vis.str = sb.String()
}

func DebugString(e Expression) string {
	vis := DebugStringVisitor{"", ""}
	e.Visit(&vis)
	return vis.str
}
