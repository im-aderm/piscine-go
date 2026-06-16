package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	arrOfDigits := []int{}
	for n > 0 {
		last := n % 10
		arrOfDigits = append(arrOfDigits, last)
		n = n / 10
	}

	for i := len(arrOfDigits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(arrOfDigits[i] + '0'))
	}

}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
