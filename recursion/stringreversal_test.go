package recursion

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	t.Run("Test String Reversal", func(t *testing.T) {
		got := ReverseString("Ismail")
		want := "liamsI"

		assertCorrectMesaage(t, got, want)
	})
}

func assertCorrectMesaage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func ExampleReverseString() {
	fmt.Println(ReverseString("Ismail"))

	// output:
	// liamsI
}

func BenchmarkReverseString(b *testing.B) {
	for b.Loop() {
		ReverseString("Ismail")
	}
}
