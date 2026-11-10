package problems

func SubArraySum(array []int, k int) int {
	if len(array) < k {
		return 0
	}

	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += array[i]
	}
	maximum := windowSum
	for i := k; i < len(array); i++ {
		windowSum += array[i] - array[i-k]
		if windowSum > maximum {
			maximum = windowSum
		}
	}
	return maximum
}
