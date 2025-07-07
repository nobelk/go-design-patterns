package workerpool

import (
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	Id           int
	RandomNumber int
}

type Result struct {
	Job                   Job
	SumOfDigits, WorkerId int
}

func digits(number int) int {
	sum := 0
	n := number
	for n != 0 {
		digit := n % 10
		sum += digit
		n /= 10
	}
	time.Sleep(5 * time.Second)
	return sum
}

func CreateWorkerPool(numWorkers int, jobs chan Job, results chan Result) {
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for job := range jobs {
				sum := Result{job, digits(job.RandomNumber), i}
				results <- sum
			}
		}(i)
	}
	wg.Wait()
	close(results)
}

func Allocate(noOfJobs int, jobs chan Job) {
	for i := 0; i < noOfJobs; i++ {
		randomNumber := rand.Intn(1000000) + 1
		job := Job{i, randomNumber}
		jobs <- job
	}
	close(jobs)
}
