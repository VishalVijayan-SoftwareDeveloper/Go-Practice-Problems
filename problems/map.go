package problems

import (
	"fmt"
	"strings"
)

func CountWords(input string) map[string]int {

	words := strings.Fields(strings.ToLower(input))
	count := make(map[string]int)
	for _, val := range words {
		count[val]++
	}
	fmt.Println(count)
	return count
}
