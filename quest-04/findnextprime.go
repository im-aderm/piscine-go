package main

import "fmt"

func FindNextPrime(n int) int {
	for {
		if ssPrime(n) {
			return n
		}
		n++
	}
}

func ssPrime(n int) bool {
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
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}
