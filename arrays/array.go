package arrays

func Sum(inp []int) int {
	output := 0
	for _, val := range inp {
		output += val
	}
	return output
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}
	return sums
}


func SumTail(tailsToSum ...[]int) []int {
	var sums []int
	for _, slice := range tailsToSum {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
