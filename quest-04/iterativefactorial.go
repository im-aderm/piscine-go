package main

import "fmt"

func IterativeFactorial(n int) int {
	if n < 0 {
		return 0
	}

	fact := 1
	for n > 0 {
		fact = fact * n
		n--
	}

	return fact
}

// optimised version

func iterativeFact(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	arg := 120
	fmt.Println(IterativeFactorial(arg))
}
