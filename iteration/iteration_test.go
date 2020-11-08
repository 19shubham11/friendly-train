package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat5Times(t *testing.T) {
	expected := Repeat5Times("a")
	actual := "aaaaa"

	if expected != actual {
		t.Errorf("Expected - %q, Actual - %q", expected, actual)
	}
}


func TestRepeatNTimes(t *testing.T) {
	expected := RepeatNTimes("a", 7)
	actual := "aaaaaaa"

	if expected != actual {
		t.Errorf("Expected - %q, Actual - %q", expected, actual)
	}
}

func BenchmarkRepeat5Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
        Repeat5Times("a")
    }
}

func BenchmarkRepeatNTimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
        RepeatNTimes("a", 7)
    }
}

func ExampleRepeatNTimes() {
	expected := RepeatNTimes("c", 6)
	fmt.Println(expected)
	// Output: cccccc
}
