package main

import "fmt"

func BasicAtoi(s string) int {
	if s == "" {
		return 0
	}

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
	fmt.Println(BasicAtoi("agsj"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
