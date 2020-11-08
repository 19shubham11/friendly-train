package arrays

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {

	t.Run("with size x", func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5, 7}
		got := Sum(numbers)
		want := 22
	
		if got != want {
			t.Errorf("Got - %d, want %d, given %v", got, want, numbers)
		}
	})

	t.Run("with empty slices", func(t *testing.T){
		numbers := []int{}
		got := Sum(numbers)
		want := 0
	
		if got != want {
			t.Errorf("Got - %d, want %d, given %v", got, want, numbers)
		}
	})
}


func TestSumAll(t *testing.T) {
	t.Run("Get sum of the given slices", func(t *testing.T){
		got := SumAll([]int{1, 2}, []int{0, 9}, []int{19, 11})
		want := []int{3, 9, 30}

		if !reflect.DeepEqual(got, want) {
				t.Errorf("Got - %v, want %v", got, want)
			}
	})
}

func TestSumTail(t* testing.T) {
	t.Run("Get the sum of tails of the given slices", func(t *testing.T) {
		got := SumTail([]int{1, 2}, []int{0, 9}, []int{19, 11, 93})
		want := []int{2, 9, 104}

		if !reflect.DeepEqual(got, want) {
				t.Errorf("Got - %v, want %v", got, want)
			}
	})

		t.Run("Get the sum of tails of the given slices", func(t *testing.T) {
		got := SumTail([]int{1, 2}, []int{0, 9}, []int{})
		want := []int{2, 9, 0}

		if !reflect.DeepEqual(got, want) {
				t.Errorf("Got - %v, want %v", got, want)
			}
	})
}
