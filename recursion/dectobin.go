package recursion

func DecToBin(dec int, result string) string {
	if dec == 0 {
		if result == "" {
			return "0"
		}
		return result
	}

	// Prepend the remainder (add to front, not back)
	result = string(rune('0'+dec%2)) + result
	return DecToBin(dec/2, result)
}
