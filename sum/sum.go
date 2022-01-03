package sum

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
