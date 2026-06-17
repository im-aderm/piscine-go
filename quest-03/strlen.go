package main

import "fmt"

func StrLen(s string) int {
	return len(s)
}

func strLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func main() {
	l := strLen("Hello World!")
	fmt.Println(l)
}
