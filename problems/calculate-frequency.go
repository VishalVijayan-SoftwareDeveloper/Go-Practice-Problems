package problems

import "fmt"

func FrequencyCalculator(Heartbeat []string) (string, int) {

	freq := make(map[string]int)
	maxFreq := 0
	maxKey := ""
	for _, beat := range Heartbeat {

		freq[beat]++
		if freq[beat] > maxFreq {
			maxFreq = freq[beat]
			maxKey = beat
		}
	}
	fmt.Printf("Most frequent: %s (%d)\n", maxKey, maxFreq)
	return maxKey, maxFreq
}
