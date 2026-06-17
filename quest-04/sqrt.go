package main

import "fmt"

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}

	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}

	return 0
}

func funcSqrtBinarySearch(nb int) int {
	if nb <= 0 {
		return 0
	}

	low := 1
	high := nb

	for low <= high {
		mid := (low + high) / 2
		square := mid * mid

		if square == nb {
			return mid
		}

		if square < nb {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return 0
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
