package main

import (
	"fmt"
)

func SortIntegerTable(table []int) {
	for i := range table {
		minIdx := i
		for j := i + 1; j < len(table); j++ {
			if table[j] < table[minIdx] {
				minIdx = j
			}
		}
		table[i], table[minIdx] = table[minIdx], table[i]
	}
}

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}
