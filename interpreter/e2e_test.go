package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/priyendra/dlisp/expression"
)

// A program is a list of strings
type Program []string

func Eval(prog Program) (expression.Expression, error) {
	env := NewStdEnvironment()
	var result, parsed expression.Expression
	var err error
	for _, exp := range prog {
		fmt.Println(exp)
		parsed, err = expression.Parse(exp)
		if err != nil {
			return nil, err
		}
		result, err = expression.Eval(parsed, env)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func checkValue(
	result expression.Expression,
	expectedType expression.Type,
	value interface{}) bool {
	if expression.ToType(result) != expectedType {
		return false
	}
	switch expectedType {
	case expression.BOOL:
		return expression.AsBool(result) == value.(bool)
	case expression.INT:
		return expression.AsInt(result) == value.(int64)
	case expression.FLOAT:
		return expression.AsFloat(result) == value.(float64)
	case expression.SYMBOL:
		return expression.AsSymbol(result) == value.(string)
	case expression.FUNCTION:
		// cannot do much for matching FUNCTION
		return true
	case expression.LIST:
		// cannot do much for matching LIST
		return true
	}
	return false
}

func TestOne(t *testing.T) {
	result, err := Eval(Program{
		"(define r 10)",
		"(* pi (* r r))",
	})
	if err != nil {
		t.Error("Could not eval program", err)
	}
	if !checkValue(result, expression.FLOAT, math.Pi*100) {
		t.Error("Expected 314.159265f, got ", expression.DebugString(result))
	}
}

func TestTwo(t *testing.T) {
	result, err := Eval(Program{
		"(if (> (* 11 11) 120) (* 7 6) oops)",
	})
	if err != nil {
		t.Error("Could not eval program", err)
	}
	if !checkValue(result, expression.INT, int64(42)) {
		t.Error("Expected 42, got ", expression.DebugString(result))
	}
}

func TestThree(t *testing.T) {
	result, err := Eval(Program{
		"(define circle-area (lambda (r) (* pi (* r r))))",
		"(circle-area 10)",
	})
	if err != nil {
		t.Error("Could not eval program", err)
	}
	if !checkValue(result, expression.FLOAT, math.Pi*100) {
		t.Error("Expected 314.159265f, got ", expression.DebugString(result))
	}
}

func TestFour(t *testing.T) {
	result, err := Eval(Program{
		"(define fact (lambda (n) (if (<= n 1) 1 (* n (fact (- n 1))))))",
		"(fact 10)",
	})
	if err != nil {
		t.Error("Could not eval program", err)
	}
	if !checkValue(result, expression.INT, int64(3628800)) {
		t.Error("Expected 3628800, got ", expression.DebugString(result))
	}
}

func TestFive(t *testing.T) {
	result, err := Eval(Program{
		"(define first car)",
		"(define rest cdr)",
		"(define count1st (lambda (item L) " +
			"(if (> (len L) 0) " +
			"(+ (if (= item (first L)) 1 0) (count1st item (rest L))) 0)))",
		"(count1st 0 (quote (0 1 2 3 0 0)))",
	})
	if err != nil {
		t.Error("Could not eval program", err)
	}
	if !checkValue(result, expression.INT, int64(3)) {
		t.Error("Expected 3, got ", expression.DebugString(result))
	}
}
