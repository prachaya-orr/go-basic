package main

import "fmt"

// jobs
type number struct {
	a int
	b int
}

// results
type sum struct {
	r int
}

func workers(jobsCh <-chan number, resultsCh chan<- sum) {
	for job := range jobsCh {
		resultsCh <- sum{r: summation(job.a, job.b)}
	}
}

func summation(a, b int) int {
	return a + b
}

func main() {
	// Worker Pool
	nums := []number{
		{a: 1, b: 2},
		{a: 2, b: 3},
		{a: 3, b: 4},
		{a: 4, b: 5},
		{a: 5, b: 6},
		{a: 1, b: 2},
		{a: 2, b: 3},
		{a: 3, b: 4},
		{a: 4, b: 5},
		{a: 5, b: 6},
	}

	jobsCh := make(chan number, len(nums))
	resultsCh := make(chan sum, len(nums))

	for _, n := range nums {
		jobsCh <- n
	}
	close(jobsCh)

	numberWorkers := 2
	for w := 0; w < numberWorkers; w++ {
		// go routines
		go workers(jobsCh, resultsCh)
	}

	results := make([]sum, 0)
	for a := 0; a < len(nums); a++ {
		temp := <-resultsCh
		results = append(results, temp)
	}

	fmt.Println(results)

}
