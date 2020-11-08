package numbers

import (
	"testing"
	"fmt"
)


func TestAdd(t *testing.T) {
	expected := Sum(2,3)
	actual := 5

	if expected != actual {
		t.Errorf("Expected - %d, Actual - %d", expected, actual)
	}
}



func ExampleSum() {
	sum := Sum(4,4)
	fmt.Println(sum)

	 // Output: 8
}