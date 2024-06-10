package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Scanln(&num1)
	fmt.Scanln(&operator)
	fmt.Scanln(&num2)
	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}
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
