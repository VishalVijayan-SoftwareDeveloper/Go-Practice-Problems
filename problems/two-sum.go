package problems

func TwoSum(numbers []int, target int) []int {
	seen := make(map[int]int)

	for i, number := range numbers {

		complement := target - number

		if id, exist := seen[complement]; exist {
			return []int{id, i}
		}
		seen[number] = i
	}
	return nil
}
