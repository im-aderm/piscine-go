package main

import "fmt"

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrime(11))
	fmt.Println(IsPrime(4))
}
