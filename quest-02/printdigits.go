package main

import "github.com/01-edu/z01"

func main() {
	start, end := '0', '9'
	for digit := start; digit <= end; digit++ {
		z01.PrintRune(digit)
	}

	z01.PrintRune('\n')
}

