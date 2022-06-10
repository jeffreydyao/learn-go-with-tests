package iteration

import (
	"fmt"
	"strings"
	"testing"
)

const repeatCount = 5

func ExampleRepeat() {
	repeated := strings.Repeat("a", repeatCount)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", repeatCount)
	}
}

func TestRepeat(t *testing.T) {
	repeated := strings.Repeat("a", repeatCount)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
