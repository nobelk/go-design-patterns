package producerconsumer

import (
	"math/rand"
	"time"
)

func Increment(previous int) int {
	time.Sleep(time.Duration(rand.Intn(100)+1) * time.Millisecond)
	return previous + 1
}
