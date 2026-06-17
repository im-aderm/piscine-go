package recursion

import "testing"

func TestDecToBin(t *testing.T) {
	t.Run("Test Dec To Bin", func(t *testing.T) {
		got := DecToBin(65, "")
		want := "1000001"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
