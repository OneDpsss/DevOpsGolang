package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(1, 2)
	expected := 3.0
	if result != expected {
		t.Errorf("add(1, 2) = %f; expected %f", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5, 3)
	expected := 2.0
	if result != expected {
		t.Errorf("subtract(5, 3) = %f; expected %f", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(2, 3)
	expected := 6.0
	if result != expected {
		t.Errorf("multiply(2, 3) = %f; expected %f", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result, err := divide(6, 2)
	expected := 3.0
	if err != nil || result != expected {
		t.Errorf("divide(6, 2) = %f, %v; expected %f, nil", result, err, expected)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(6, 0)
	if err == nil {
		t.Errorf("divide(6, 0) did not return an error")
	}
}

func TestCalculateAdd(t *testing.T) {
	result, err := calculate(1, 2, "+")
	expected := 3.0
	if err != nil || result != expected {
		t.Errorf("calculate(1, 2, '+') = %f, %v; expected %f, nil", result, err, expected)
	}
}

func TestCalculateSubtract(t *testing.T) {
	result, err := calculate(5, 3, "-")
	expected := 2.0
	if err != nil || result != expected {
		t.Errorf("calculate(5, 3, '-') = %f, %v; expected %f, nil", result, err, expected)
	}
}

func TestCalculateMultiply(t *testing.T) {
	result, err := calculate(2, 3, "*")
	expected := 6.0
	if err != nil || result != expected {
		t.Errorf("calculate(2, 3, '*') = %f, %v; expected %f, nil", result, err, expected)
	}
}

func TestCalculateDivide(t *testing.T) {
	result, err := calculate(6, 2, "/")
	expected := 3.0
	if err != nil || result != expected {
		t.Errorf("calculate(6, 2, '/') = %f, %v; expected %f, nil", result, err, expected)
	}
}

func TestCalculateDivideByZero(t *testing.T) {
	_, err := calculate(6, 0, "/")
	if err == nil {
		t.Errorf("calculate(6, 0, '/') did not return an error")
	}
}

func TestCalculateInvalidOperator(t *testing.T) {
	_, err := calculate(6, 2, "^")
	if err == nil {
		t.Errorf("calculate(6, 2, '^') did not return an error for invalid operator")
	}
}
