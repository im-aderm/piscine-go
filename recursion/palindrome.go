package recursion

func IsPalindrome(s string) bool {

	// edge cases

	if len(s) == 0 || len(s) == 1 {
		return true
	}
	if s == ReverseString2(s) {
		return true
	}

	return false
}

// another version
func IsPalindrome2(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	if s[0] == s[len(s)-1] {
		return IsPalindrome2(s[1 : len(s)-1])
	}

	return false
}
