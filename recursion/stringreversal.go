package recursion

// using recursion
func ReverseString(s string) string {
	if s == "" {
		return ""
	}
	return ReverseString(s[1:] + string(s[0]))
}

// Using rune conversion
func ReverseString2(s string) string {
	runes := []rune(s)
	return reveseRunes(runes)

}

func reveseRunes(r []rune) string {
	if len(r) == 0 {
		return ""
	}

	return reveseRunes(r[1:]) + string(r[0])
}
