package main

import "fmt"

func DivMod(a, b int, div, mod *int) {
	*div = a / b
	*mod = a % b
}

func main() {
	a := 13
	b := 2

	var div int
	var mod int

	DivMod(a, b, &div, &mod)
	fmt.Println(a)
	fmt.Println(b)
}
