package parser

import (
	"errors"
	"fmt"
	"strings"
)

type tokenStream struct {
	tokens []string
}

func (s *tokenStream) peek() string {
	return s.tokens[0]
}

func (s *tokenStream) next() string {
	answer := s.tokens[0]
	s.tokens = s.tokens[1:]
	return answer
}

func (s *tokenStream) done() bool {
	return len(s.tokens) == 0
}

type AstNodeType int

const (
	AST_SYMBOL = iota
	AST_INT_LITERAL
	AST_LIST
)

type AstNode struct {
	astType  AstNodeType
	symbol   string
	literal  int // TODO: Support literals other than just integers.
	children []*AstNode
}

func (node *AstNode) Type() AstNodeType {
	return node.astType
}

func (node *AstNode) Symbol() string {
	return node.symbol
}

func (node *AstNode) Literal() int {
	return node.literal
}

func (node *AstNode) NumChildren() int {
	return len(node.children)
}

func (node *AstNode) Child(i int) *AstNode {
	return node.children[i]
}

func (node *AstNode) debugStringHelper(indentAmt int) string {
	var sb strings.Builder
	indent := strings.Repeat(" ", indentAmt)
	switch node.astType {
	case AST_SYMBOL:
		sb.WriteString(fmt.Sprintf("%sSym %s", indent, node.symbol))
	case AST_INT_LITERAL:
		sb.WriteString(fmt.Sprintf("%sLit %d", indent, node.literal))
	case AST_LIST:
		sb.WriteString(indent)
		sb.WriteString("List {\n")
		for _, child := range node.children {
			sb.WriteString(child.debugStringHelper(indentAmt + 2))
			sb.WriteString("\n")
		}
		sb.WriteString(indent)
		sb.WriteString("}")
	}
	return sb.String()
}

func (node *AstNode) DebugString() string {
	return node.debugStringHelper(0)
}

// TODO: Improve this is a very basic tokenizer. Does not handle quotes, does
// not handle multiline etc.
func tokenize(input string) tokenStream {
	input = strings.Replace(input, "(", " ( ", -1)
	input = strings.Replace(input, ")", " ) ", -1)
	return tokenStream{strings.Fields(input)}
}

func parseAtom(token string) (AstNode, error) {
	node := AstNode{}
	if count, _ := fmt.Sscanf(token, "%d", &node.literal); count == 1 {
		node.astType = AST_INT_LITERAL
	} else {
		node.symbol = token
		node.astType = AST_SYMBOL
	}
	return node, nil
}

func parseList(tokens *tokenStream) (AstNode, error) {
	node := AstNode{}
	node.astType = AST_LIST
	for tokens.peek() != ")" && !tokens.done() {
		child, err := parseHelper(tokens)
		if err != nil {
			return node, err
		}
		node.children = append(node.children, &child)
	}
	if tokens.peek() != ")" {
		return AstNode{}, errors.New("Unexpected syntax error")
	} else {
		tokens.next() // remove the trailing ')'
	}
	return node, nil
}

func parseHelper(tokens *tokenStream) (AstNode, error) {
	astNode := AstNode{}
	if tokens.done() {
		return astNode, errors.New("Unexpected EOF")
	}
	next := tokens.next()
	if next == "(" {
		return parseList(tokens)
	} else if next == ")" {
		return astNode, errors.New("Unexpected syntax error")
	}
	return parseAtom(next)
}

func Parse(input string) (AstNode, error) {
	tokens := tokenize(input)
	node, err := parseHelper(&tokens)
	if !tokens.done() {
		return AstNode{}, errors.New("Excess unparsed input")
	}
	return node, err
}
