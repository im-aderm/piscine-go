package main

import "fmt"

func PrintOne(n *int) {
	*n = 1
}

func main() {
	var x int
	PrintOne(&x)
	fmt.Println(x)
}
