package main

import (
	"Go-Practice-Problems/problems"
	"fmt"
)

func main() {

	//Find duplicate string
	devices := []string{
		"device-1",
		"device-2",
		"device-3",
		"device-1",
		"device-4",
		"device-2",
		"device-5",
	}
	//problems.FindDuplicates(devices)

	res := problems.BetterFindDuplicateOneMap(devices)
	fmt.Println("Duplicates devices are:", res)

	heartbeats := []string{
		"device-A",
		"device-B",
		"device-A",
		"device-C",
		"device-B",
		"device-A",
	}
	problems.FrequencyCalculator(heartbeats)

	//Two Sum

	number := []int{2, 8, 1, 7, 15}
	target := 9
	problems.TwoSum(number, target)
	//fmt.Println(sum)

	/*	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}*/
	intervals := [][]int{
		{1, 10},
		{2, 3},
		{4, 5},
		{11, 15},
	}

	problems.MergeOverlap(intervals)
}
