package fanout

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Number = struct {
	squared, wId int
}

func GeneratePipeline(numbers []int) <-chan int {
	outCh := make(chan int)
	go func() {
		for _, n := range numbers {
			outCh <- n
		}
		close(outCh)
	}()
	return outCh
}

func squareNumber(in <-chan int, val int) <-chan Number {
	outCh := make(chan Number)
	go func(val int) {
		for n := range in {
			outCh <- Number{
				n * n,
				val,
			}
		}
		close(outCh)
	}(val)
	return outCh
}

func displayData(cs ...<-chan Number) {
	for _, val := range cs {
		wg.Add(1)
		go func(value <-chan Number) {
			defer wg.Done()
			for val := range value {
				fmt.Printf("The squared number is %d\n and wId is %d\n", val.squared, val.wId)
			}
		}(val)
	}
}

func RunFanout() {
	randomNumbers := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	inputCh := GeneratePipeline(randomNumbers)
	c1 := squareNumber(inputCh, 1)
	c2 := squareNumber(inputCh, 2)
	c3 := squareNumber(inputCh, 3)
	displayData(c1, c2, c3)
	wg.Wait()
}
