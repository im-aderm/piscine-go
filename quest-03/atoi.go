package main

import "fmt"

func Atoi(s string) int {

	if s == "" {
		return 0
	}

	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	if start > 0 {
		if len(s) == 1 {
			return 0
		}

		if s[start] == '-' || s[start] == '+' {
			return 0
		}

	}

	result := 0
	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			digit := int(ch - '0')
			result = result*10 + digit
		} else {
			return 0
		}
	}

	return result * sign
}

func main() {

	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
