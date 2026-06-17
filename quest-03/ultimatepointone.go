package main

import "fmt"

func UltimatePrintOne(n ***int) {
	***n = 1
}

func main() {
	a := 0
	b := &a
	c := &b

	UltimatePrintOne(&c)
	fmt.Println(c)
}
