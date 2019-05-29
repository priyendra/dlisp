package expression

import (
	"errors"
	"strconv"
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

// TODO: Improve this is a very basic tokenizer. Does not handle quotes, does
// not handle multiline etc.
func tokenize(input string) tokenStream {
	input = strings.Replace(input, "(", " ( ", -1)
	input = strings.Replace(input, ")", " ) ", -1)
	return tokenStream{strings.Fields(input)}
}

func parseAtom(token string) (Expression, error) {
	// TODO(deshwal): Validate the tokens via regular expressions
	if i, err := strconv.ParseInt(token, 10, 64); err == nil {
		return IntLiteral(i), nil
	} else if f, err := strconv.ParseFloat(token, 64); err == nil {
		return FloatLiteral(f), nil
	}
	return Symbol(token), nil
}

func parseList(tokens *tokenStream) (Expression, error) {
	c := Compound{}
	for tokens.peek() != ")" && !tokens.done() {
		child, err := parseHelper(tokens)
		if err != nil {
			return nil, err
		}
		c = append(c, child)
	}
	if tokens.peek() != ")" {
		return nil, errors.New("Unexpected syntax error")
	} else {
		tokens.next() // remove the trailing ')'
	}
	return c, nil
}

func parseHelper(tokens *tokenStream) (Expression, error) {
	if tokens.done() {
		return nil, errors.New("Unexpected EOF")
	}
	next := tokens.next()
	if next == "(" {
		return parseList(tokens)
	} else if next == ")" {
		return nil, errors.New("Unexpected syntax error")
	}
	return parseAtom(next)
}

func Parse(input string) (Expression, error) {
	tokens := tokenize(input)
	expr, err := parseHelper(&tokens)
	if !tokens.done() {
		return nil, errors.New("Excess unparsed input")
	}
	return expr, err
}
