package utils

func Abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Factorial(number int) int {
	// if the number has reached 1 then we have to
	// return 1 as 1 is the minimum value we have to multiply with
	if number <= 1 {
		return 1
	}

	// multiplying with the current number and calling the function
	// for 1 lesser number
	factorialOfNumber := number * Factorial(number-1)

	// return the factorial of the current number
	return factorialOfNumber
}

func Pow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
