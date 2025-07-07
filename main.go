package main

import (
	"fmt"
	"sync"

	"github.com/nobelk/go-design-patterns/generator"
	"github.com/nobelk/go-design-patterns/pipeline"
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
}
