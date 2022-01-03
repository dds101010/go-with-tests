package sum

import (
	"reflect"
	"testing"
)

func testSumInt(t testing.TB, input interface{}, got, want interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v, input: %v", got, want, input)
	}
}

func TestSumFunc(t *testing.T) {
	t.Run("Sum of array", func(t *testing.T) {
		arr := [...]int{1, 2, 3, 4, 5}
		// arr := [5]int{1, 2, 3, 4, 5}

		want := 15
		got := Sum(arr)

		testSumInt(t, arr, got, want)
	})

	t.Run("Sum of slice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}

		want := 10
		got := SumSlice(slice)

		testSumInt(t, slice, got, want)
	})

	t.Run("SumAll of slices", func(t *testing.T) {
		// []int in the internal can be omitted
		slices := [][]int{{1, 2, 3}, {4, 5, 6}}

		want := []int{6, 15}
		got := SumAll(slices[0], slices[1])

		testSumInt(t, slices, got, want)
	})
}

func TestSliceTail(t *testing.T) {
	t.Run("Sum of Tails Basic", func(t *testing.T) {
		slices := [][]int{{2, 9}, {6, 3}}

		// spread operator
		got := SumAllTails(slices...)
		want := []int{9, 3}

		testSumInt(t, slices, got, want)
	})

	t.Run("Sum of Tails Complex", func(t *testing.T) {
		slices := [][]int{{}, {6, 3}}

		// spread operator
		got := SumAllTails(slices...)
		want := []int{0, 3}

		testSumInt(t, slices, got, want)
	})
}
