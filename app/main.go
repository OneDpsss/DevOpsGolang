package main

import "fmt"

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return add(num1, num2), nil
	case "-":
		return subtract(num1, num2), nil
	case "*":
		return multiply(num1, num2), nil
	case "/":
		return divide(num1, num2)
	default:
		return 0, fmt.Errorf("неверный оператор")
	}
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return a / b, nil
}
