package pipeline

import (
	"fmt"
	"sync"
)

func DisplayData(inCh <-chan Operation) <-chan string {
	outCh := make(chan string)
	go func() {
		wg := &sync.WaitGroup{}
		for input := range inCh {
			wg.Add(1)
			go concatenateValue(input, outCh, wg)
		}
		wg.Wait()
		close(outCh)
	}()
	return outCh
}

func concatenateValue(input Operation, outCh chan string, wg *sync.WaitGroup) {
	concat := fmt.Sprintf("ID: %d, Multiply: %d, Addition: %d", input.id, input.multiply, input.addition)
	outCh <- concat
	wg.Done()
}
