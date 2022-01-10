package sum

import (
	"fmt"
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

func TestSliceInternals1(t *testing.T) {
	// slice "a" with len = 1 million
	a := make([]int, 1e6)

	// even though "b" len = 2, it points to the same underlying array "a" points to
	b := a[:2]

	// create a copy of the slice so "a" can be garbage collected
	c := make([]int, len(b))
	copy(c, b)
	fmt.Println(c)
}

func TestSliceInternals2(t *testing.T) {
	x := [3]string{"Dwight", "Angela", "Jimothy"}

	// slice "y" points to the underlying array "x"
	y := x[:]

	z := make([]string, len(x))
	// slice "z" is a copy of the slice created from array "x"
	copy(z, x[:])

	// the value at index 1 is now "Robert" for both "y" and "x"
	y[1] = "Robert"

	fmt.Printf("x - %T %v\n", x, x)
	fmt.Printf("y - %T %v\n", y, y)
	fmt.Printf("z - %T %v\n", z, z)
}
