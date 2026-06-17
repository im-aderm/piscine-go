package main

import "fmt"

func IterativePower(nb int, power int) int {
	result := 1

	if nb == 0 {
		return 0
	}

	if power < 0 {
		return 0
	}

	if nb == 0 && power == 0 {
		return 1
	}
	for i := 0; i < power; i++ {
		result *= nb
	}

	return result
}

func main() {
	fmt.Println(IterativePower(4, 3))
}
