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

func TestSix(t *testing.T) {
	result, err := Eval(Program{
		"(define twice (lambda (x) (* 2 x)))",
		"(define repeat (lambda (f) (lambda (x) (f (f x)))))",
		"((repeat twice) 10)",
	})
	if err != nil {
		t.Error("Could not eval program", err)
	}
	if !checkValue(result, expression.INT, int64(40)) {
		t.Error("Expected 40, got ", expression.DebugString(result))
	}
}

func TestSeven(t *testing.T) {
	result, err := Eval(Program{
		"(define twice (lambda (x) (* 2 x)))",
		"(define repeat (lambda (f) (lambda (x) (f (f x)))))",
		"((repeat (repeat twice)) 10)",
	})
	if err != nil {
		t.Error("Could not eval program", err)
	}
	if !checkValue(result, expression.INT, int64(160)) {
		t.Error("Expected 160, got ", expression.DebugString(result))
	}
}

func TestEight(t *testing.T) {
	result, err := Eval(Program{
		"(define fib (lambda (n) (if (< n 2) 1 (+ (fib (- n 1)) (fib (- n 2))))))",
		"(define range (lambda (a b) " +
			"(if (>= a b) (quote ()) (append (range a (- b 1)) (- b 1)))))",
		"(map fib (range 0 10))",
	})
	if err != nil {
		t.Error("Could not eval program", err)
	}
	if expression.ToType(result) != expression.LIST {
		t.Error("Expected result type LIST")
	}
	list := expression.AsList(result)
	if !checkValue(list[0], expression.INT, int64(1)) {
		t.Error("Expected 1, got ", expression.DebugString(result))
	}
	if !checkValue(list[1], expression.INT, int64(1)) {
		t.Error("Expected 1, got ", expression.DebugString(result))
	}
	if !checkValue(list[2], expression.INT, int64(2)) {
		t.Error("Expected 2, got ", expression.DebugString(result))
	}
	if !checkValue(list[3], expression.INT, int64(3)) {
		t.Error("Expected 3, got ", expression.DebugString(result))
	}
	if !checkValue(list[4], expression.INT, int64(5)) {
		t.Error("Expected 5, got ", expression.DebugString(result))
	}
	if !checkValue(list[5], expression.INT, int64(8)) {
		t.Error("Expected 8, got ", expression.DebugString(result))
	}
	if !checkValue(list[6], expression.INT, int64(13)) {
		t.Error("Expected 13, got ", expression.DebugString(result))
	}
	if !checkValue(list[7], expression.INT, int64(21)) {
		t.Error("Expected 21, got ", expression.DebugString(result))
	}
	if !checkValue(list[8], expression.INT, int64(34)) {
		t.Error("Expected 34, got ", expression.DebugString(result))
	}
	if !checkValue(list[9], expression.INT, int64(55)) {
		t.Error("Expected 55, got ", expression.DebugString(result))
	}
}
