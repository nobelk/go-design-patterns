package pipeline

type Operation struct {
	id       int64
	multiply int64
	addition int64
}

func PrepareData(inCh <-chan int64) <-chan Operation {
	outCh := make(chan Operation)
	go func() {
		for id := range inCh {
			input := Operation{id: id, multiply: id * 2, addition: id + 5}
			outCh <- input
		}
		close(outCh)
	}()
	return outCh
}
