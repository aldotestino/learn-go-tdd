package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectRepeat := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected %q but got %q", got, want)
		}

	}

	t.Run("Repeating N times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		assertCorrectRepeat(t, repeated, expected)
	})

	t.Run("Repeating N times with String package", func(t *testing.T) {
		repeated := RepeatWithStrings("b", 5)
		expected := "bbbbb"

		assertCorrectRepeat(t, repeated, expected)
	})

}

func ExampleRepeat() {
	repeated := Repeat("a", 10)
	fmt.Println(repeated)
	// Output: aaaaaaaaaa
}

func ExampleRepeatWithStrings() {
	repeated := RepeatWithStrings("b", 5)
	fmt.Println(repeated)
	// Output: bbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func BenchmarkRepeatWithString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithStrings("a", 10)
	}
}
