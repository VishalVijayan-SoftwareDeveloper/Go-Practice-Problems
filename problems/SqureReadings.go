package problems

import (
	"fmt"
	"sync"
)

func workerpool(jobs <-chan Job, res chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		res <- Result{
			Index: j.Index,
			Value: j.Value * j.Value,
		}
	}
}

func WorkerSquareReadings(input []int, workerCount int) []int {
	jobs := make(chan Job, len(input))
	result := make(chan Result, workerCount)
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go workerpool(jobs, result, &wg)
	}

	for j := 0; j < len(input); j++ {
		jobs <- Job{
			Index: j,
			Value: input[j],
		}
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(result)
	}()
	res := make([]int, 0, len(input))

	for r := range result {
		res = append(res, r.Value)
	}
	fmt.Println(res)
	return res
}
