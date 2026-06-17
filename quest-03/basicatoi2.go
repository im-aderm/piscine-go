package main

import "fmt"

func BasicAtoi2(s string) int {
	result := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			digit := int(ch - '0')
			result = result*10 + digit
		} else {
			return 0
		}
	}
	return result
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
