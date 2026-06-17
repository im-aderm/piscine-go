package main

import "fmt"

func StrRev(s string) string {

	bytes := make([]byte, len(s))
	i := len(s) - 1
	j := 0
	for i >= 0 {
		bytes[j] = s[i]
		i--
		j++
	}
	return string(bytes)
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
