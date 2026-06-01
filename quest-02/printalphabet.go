package main

import "github.com/01-edu/z01"

func main() {
	start, end := 'a', 'z'
	for char := start; char <= end; char++ {
		if err := z01.PrintRune(char); err != nil {
			return
		}
	}
	z01.PrintRune('\n')
}

// alternative approach

/**
	func main() {
		char :+ 'a'
		for char <= 'z' {
			z01.PrintRune(char)
			char++
		}
	}
*/
