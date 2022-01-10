package sum

// https://go.dev/blog/slices-intro
/**
* Slicing does not copy the slice's data. It creates a new slice
* value that points to the original array. This makes slice operations
* as efficient as manipulating array indices. Therefore, modifying the
* elements (not the slice itself) of a re-slice modifies the elements
* of the original slice
*
* To increase the capacity of a slice one must create a new, larger slice
* and copy the contents of the original slice into it.
 */

/**
* An interesting property of arrays is that the size is encoded in its type.
* If you try to pass an [4]int into a function that expects [5]int, it won't compile.
 */
func Sum(arr [5]int) int {
	var sum int
	for _, num := range arr {
		sum += num
	}
	return sum
}

func SumSlice(slice []int) int {
	var sum int
	for _, num := range slice {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	res := make([]int, len(slices))

	for i, slice := range slices {
		res[i] = SumSlice(slice)
	}

	return res
}

func SumAllTails(slices ...[]int) []int {
	var res []int

	for _, slice := range slices {
		if len(slice) == 0 {
			res = append(res, 0)
		} else {
			res = append(res, SumSlice(slice[1:]))
		}
	}

	return res
}
