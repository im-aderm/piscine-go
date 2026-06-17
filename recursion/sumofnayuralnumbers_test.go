package recursion

import "testing"

func TestSumNum(t *testing.T) {
	t.Run("Test IsPalindrome", func(t *testing.T) {
		got := SumNum(10)
		want := 55

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
