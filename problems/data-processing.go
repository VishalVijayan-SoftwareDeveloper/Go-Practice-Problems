package problems

import "fmt"

type Input struct {
	Device string
	Value  int
}

func ProcessData(input []Input) map[string]int {
	seen := make(map[string]int)
	for _, r := range input {
		seen[r.Device] = seen[r.Device] + r.Value
	}
	fmt.Println(seen)
	return seen

}
