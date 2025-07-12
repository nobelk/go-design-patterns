package main

import (
	"fmt"
	"github.com/nobelk/go-design-patterns/fanin"
	"sync"

	"github.com/nobelk/go-design-patterns/fanout"
	"github.com/nobelk/go-design-patterns/generator"
	"github.com/nobelk/go-design-patterns/pipeline"
	"github.com/nobelk/go-design-patterns/producerconsumer"
	"github.com/nobelk/go-design-patterns/timeoutusingselect"
	"github.com/nobelk/go-design-patterns/workerpool"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("\n===Generator Pattern===\n")
	for num := range generator.Fibonacci(10000000) {
		fmt.Printf("Fibonacci number is: %d\n", num)
	}

	fmt.Println("\n===Pipeline Pattern===\n")
	ch := pipeline.DisplayData(pipeline.PrepareData(pipeline.GenerateData()))
	for data := range ch {
		fmt.Printf("Items: %+v\n", data)
	}

	fmt.Println("\n===Worker Pool Pattern===\n")
	jobs := make(chan workerpool.Job, 10)
	results := make(chan workerpool.Result, 10)
	numWorkers := 5
	noOfJobs := 10

	go workerpool.CreateWorkerPool(numWorkers, jobs, results)
	go workerpool.Allocate(noOfJobs, jobs)

	for i := 0; i < noOfJobs; i++ {
		result := <-results
		fmt.Printf("Job id %d, sum of digits %d, worker id %d\n", result.Job.Id, result.Job.RandomNumber, result.SumOfDigits, result.WorkerId)
	}

	fmt.Println("\n===Fanin Pool Pattern===\n")
	fanin.RunFanin()

	fmt.Println("\n===Fanout Pool Pattern===\n")
	fanout.RunFanout()

	fmt.Println("\n===Multiple Producer Multiple Consumer Pattern===\n")
	producerconsumer.RunMultiProducerMultiConsumer()

	fmt.Println("\n===Multiple Producer Single Consumer Pattern===\n")
	producerconsumer.RunMultipleProducerSingleConsumer()

	fmt.Println("\n===Single Producer Multiple Consumer Pattern===\n")
	producerconsumer.RunSingleProducerMultiConsumer()

	fmt.Println("\n===Timeout Using Select===\n")
	timeoutusingselect.RunTimeoutUsingSelect()
}
