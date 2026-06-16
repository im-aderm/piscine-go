package main

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	digits := make([]rune, n)
	for i := range n {
		digits[i] = '0' + rune(i)
	}

	for {
		for i := 0; i < n; i++ {
			z01.PrintRune(digits[i])
		}

		i := n - 1
		for i >= 0 && digits[i] == '9'-rune(n-1-i) {
			i--
		}

		if i < 0 {
			break
		}

		z01.PrintRune(',')
		z01.PrintRune(' ')

		digits[i]++
		for j := i + 1; j < n; j++ {
			digits[j] = digits[j-1] + 1
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintCombN(4)
}
