package pipeline

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func GenerateData() <-chan int64 {
	outCh := make(chan int64)
	const filePath = "integer.txt"
	go func() {
		file, _ := os.Open(filePath)
		defer close(outCh)
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimSuffix(line, "\r")
			num, _ := strconv.ParseInt(line, 10, 0)
			outCh <- num
			if err == io.EOF {
				break
			}
		}
	}()
	return outCh
}
