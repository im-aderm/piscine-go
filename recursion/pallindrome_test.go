package recursion

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Run("Test IsPalindrome", func(t *testing.T) {
		got := IsPalindrome("kayak")
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Test IsPalindrome2", func(t *testing.T) {
		got := IsPalindrome2("kayak")
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
