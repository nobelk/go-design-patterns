package generator

import (
	"fmt"
)

func Fibonacci(n int) <-chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		fmt.Println("Running Producer")
		for i, j := 0, 1; i < n; i, j = i+j, i {
			outCh <- i
		}
	}()
	return outCh
}
