package main

import (
	"fmt"

	"go-functions-deep-dive/simplemath"
)

func main() {

	numbers := []int{4, 1, 10, 4}
	a := simplemath.Sum(numbers...)
	// a := simplemath.Sum(4, 1, 10, 4)
	fmt.Printf("%d\n", a)

	answer, err := simplemath.Divide(10, 2)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}

}
