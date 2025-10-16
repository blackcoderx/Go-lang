package main

import (
	"calculator/mathops" // Corrected module-relative import path
	"fmt"
)

func calculator(op func(a, b float64) (float64, error), a, b float64) (float64, error) {
	return op(a, b)
}

func main() {

	difference, err := calculator(mathops.Subtract, 10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	}
	sum, err := calculator(mathops.Add, 10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	product, err := calculator(mathops.Multiply, 10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	division, err := calculator(mathops.Divide, 10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Difference:", difference)
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
	fmt.Println("Division:", division)

}
