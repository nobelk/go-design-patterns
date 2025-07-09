package fanin

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Number struct {
	Original, Reverse int
}

func ReadFile(filename string) (<-chan int, error) {
	outCh := make(chan int)
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Could not read file %v", err)
	}
	go func(file *os.File) {
		defer close(outCh)
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimSuffix(line, "\r")
			num, _ := strconv.Atoi(line)
			outCh <- num
			if err == io.EOF {
				break
			}
		}
	}(file)
	return outCh, nil
}

func reverseNumber(n int) int {
	result := 0

	for n > 0 {
		result = result*10 + n%10
		n /= 10

	}

	return result
}

func Merge(cs ...<-chan int) <-chan Number {
	var wg sync.WaitGroup
	out := make(chan Number)
	send := func(c <-chan int) {
		for n := range c {
			out <- Number{
				n,
				reverseNumber(n),
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
