package simplemath

import (
	"errors"
	"math"
)

type MathExpression = string

const (
	Add      = MathExpression("add")
	Subtract = MathExpression("subtract")
	Multiply = MathExpression("multiply")
)

func Sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func add(number1, number2 float64) float64 {
	return number1 + number2
}

func subtract(number1, number2 float64) float64 {
	return number1 - number2
}

func multiply(number1, number2 float64) float64 {
	return number1 * number2
}

func Divide(number1, number2 float64) (answer float64, err error) {
	if number2 == 0 {
		err = errors.New("Divide by zero error")
	}
	answer = number1 / number2
	return
}

func Expression(expression MathExpression) func(float64, float64) float64 {
	switch expression {
	case Add:
		return add
	case Subtract:
		return subtract
	case Multiply:
		return multiply
	default:
		return func(num1, num2 float64) float64 {
			return 0
		}
	}
}

func Double(num1, num2 float64, mathExpression func(float64, float64) float64) float64 {
	return 2 * mathExpression(num1, num2)
}

func PowerOfTwo() func() int64 {
	x := 1.0
	// stateful function.
	// Everytime we call this function, it add 1 to the previous value
	return func() int64 {
		x++
		return int64(math.Pow(x, 2))
	}
}
