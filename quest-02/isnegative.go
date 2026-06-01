package main

import "github.com/01-edu/z01"

func isNegative(n int) {
	if n < 0 {
		z01.PrintRune('T')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('F')
		z01.PrintRune('\n')
	}
}

func main() {
	isNegative(6)
}
