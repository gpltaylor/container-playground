package main

import (
	"fmt"
	"testing"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			results <- 0
		}
	}()

	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)

		x := 0
		if j == 3 {
			a := 10
			b := 0
			x = a / b
		}
		results <- j*2 + x
		//fmt.Println("worker", id, "finished job", j)
	}
}

func TestWorkerPool_JAGBE01(t *testing.T) {

	const numJobs = 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		answer := <-results
		fmt.Println("answer", answer)
	}
}
