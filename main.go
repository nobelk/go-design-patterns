package main

import (
	"fmt"
	"sync"

	"github.com/nobelk/go-design-patterns/fanin"
	"github.com/nobelk/go-design-patterns/generator"
	"github.com/nobelk/go-design-patterns/pipeline"
	"github.com/nobelk/go-design-patterns/producerconsumer"
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
	ch1, err := fanin.ReadFile("file1.txt")
	if err != nil {
		fmt.Println(err)
	}
	ch2, err := fanin.ReadFile("file2.txt")
	if err != nil {
		fmt.Println(err)
	}
	ch3, err := fanin.ReadFile("file3.txt")
	if err != nil {
		fmt.Println(err)
	}

	channel := fanin.Merge(ch1, ch2, ch3)

	for val := range channel {
		fmt.Println("Original number: %v Reversed number: %v",
			val.Original, val.Reverse)
	}

	fmt.Println("\n===Single Producer Single Consumer Pattern===\n")
	data := make(chan int)
	// producer
	go func() {
		defer close(data)
		for i := 0; i < 100; i++ {
			data <- producerconsumer.Increment(i)
		}
	}()

	// consumer
	for i := range data {
		fmt.Printf("Value of i: %d\n", i)
	}
}
