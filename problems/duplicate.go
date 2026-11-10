package problems

import "fmt"

func FindDuplicates(devices []string) (result []string) {
	m := make(map[string]int)
	for _, device := range devices {
		m[device]++
	}
	for device, count := range m {
		if count > 1 {
			result = append(result, device)
		}
	}
	fmt.Println("Duplicates:", result)
	return result
}

//Better version

func BetterFindDuplicate(devices []string) []string {

	seen := make(map[string]bool)
	reported := make(map[string]bool)
	var result []string

	for _, device := range devices {

		if seen[device] {
			if !reported[device] {
				result = append(result, device)
				reported[device] = true
			}
		} else {
			seen[device] = true
		}
	}
	return result
}

// Better with one map
func BetterFindDuplicateOneMap(devices []string) []string {

	seen := make(map[string]bool)
	//reported := make(map[string]bool)
	var result []string

	for _, device := range devices {
		if seen[device] {
			result = append(result, device)
			seen[device] = false
		} else if _, exists := seen[device]; !exists {
			seen[device] = true
		}
	}

	return result
}
