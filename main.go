package main

import (
	"fmt"

	"go-functions-deep-dive/simplemath"
	"go-functions-deep-dive/version"
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

	addExpression := simplemath.Expression(simplemath.Add)
	fmt.Printf("%f\n", addExpression(1, 2))

	multiplyExpression := simplemath.Expression(simplemath.Multiply)
	fmt.Printf("%f\n", multiplyExpression(3, 2))

	result := simplemath.Double(4, 6, simplemath.Expression(simplemath.Multiply)) // 2 * (4*6)
	fmt.Printf("%f\n", result)

	p2 := simplemath.PowerOfTwo() // Look into the func PowerOfTwo(), x = 1
	value := p2()                 // x = 2
	println(value)
	value = p2() // x = 3
	println(value)
	value = p2() // x = 4
	println(value)

	println("\n\nSemantic version:")
	sv := version.NewSemanticVersion(0, 1, 0)
	sv.IncrementMajor()
	sv.IncrementMajor()
	sv.IncrementMinor()
	sv.IncrementPatch()
	sv.IncrementPatch()
	sv.IncrementPatch()
	sv.IncrementPatch()
	println(sv.String())

}
