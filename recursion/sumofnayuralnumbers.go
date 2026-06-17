package recursion

func SumNum(n int) int {

	if n < 0 {
		panic("n cannot be less than 0")
	}
	if n == 0 {
		return n
	}
	return SumNum(n-1) + n
}
