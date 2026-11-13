package problems

import (
	"fmt"
	"sync"
)

type Job struct {
	Index int
	Value int
}
type Result struct {
	Index int
	Value int
}

func workerPool(jobs <-chan Job, result chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		result <- Result{
			Index: job.Index,
			Value: job.Value * job.Value,
		}
	}
}
func SquareReadings(readings []int, workerCount int) []int {

	jobs := make(chan Job, len(readings))
	results := make(chan Result, workerCount)
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go workerPool(jobs, results, &wg)
	}

	for j := 0; j < len(readings); j++ {
		jobs <- Job{
			Index: j,
			Value: readings[j],
		}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()
	squared := make([]int, 0, len(readings))

	for result := range results {
		squared = append(squared, result.Value)
	}
	fmt.Println(squared)
	return squared
}
