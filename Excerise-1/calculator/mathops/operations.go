package mathops

import (
	"errors"
)

func Add(a, b float64) (float64, error) {
	return a + b, nil
}

func Multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func Subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func Divide(a, b float64) (float64, error) {
	// Validate input before performing the division.
	if b == 0 {
		return 0, errors.New("really, you cannot divide by zero")
	}
	return a / b, nil
}

// The `validatediv` function had syntax errors and incorrect logic.
// It's better to handle validation directly inside the `Divide` function
// for clarity, so I have removed it.
