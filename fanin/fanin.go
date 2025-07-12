package fanin

import (
	"bufio"
	"fmt"
	"github.com/nobelk/go-design-patterns/producerconsumer"
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

func RunFanin() {
	ch1, err := ReadFile("file1.txt")
	if err != nil {
		fmt.Println(err)
	}
	ch2, err := ReadFile("file2.txt")
	if err != nil {
		fmt.Println(err)
	}
	ch3, err := ReadFile("file3.txt")
	if err != nil {
		fmt.Println(err)
	}

	channel := Merge(ch1, ch2, ch3)

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
