package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeating a character 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})

	t.Run("counting the repeated characters", func(t *testing.T) {
		repeatedString := Repeat("a", 100)
		repeatedCount := strings.Count(repeatedString, "a")
		expected := 100

		if repeatedCount != expected {
			t.Errorf("expected '%d' but got '%d'", expected, repeatedCount)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 4)
	fmt.Printf("%q", repeated)
	// Output: "aaaa"
}
