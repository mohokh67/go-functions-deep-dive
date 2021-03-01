package simplemath

import (
	"errors"
)

func Sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func Divide(number1, number2 float64) (answer float64, err error) {
	if number2 == 0 {
		err = errors.New("Divide by zero error")
	}
	answer = number1 / number2
	return
}
