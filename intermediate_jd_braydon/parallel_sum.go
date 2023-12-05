// JD Waldron and Braydon Hughes
// Program that sums up an array concurrently
package main

import (
	"fmt"
	"sync"
)

var (
	numArray [10000]int
	totalSum int
	mu       sync.Mutex
	wg       sync.WaitGroup
	N        int
)

func work(low int, high int) {
	var localSum int
	mu.Lock()
	for i := low; i < high; i++ {
		localSum += numArray[i]
	}
	totalSum += localSum
	mu.Unlock()
	wg.Done()
}

func main() {
	numThreads := 10
	totalSum = 0
	N = len(numArray)

	i := 0
	for i < N {
		numArray[i] = i + 1
		i++
	}

	workNum := N / numThreads

	for i := 0; i < numThreads; i++ {
		low := i * workNum
		high := (i + 1) * workNum
		wg.Add(1)
		go work(low, high)
	}

	wg.Wait()
	fmt.Printf("SeqSum: %d, ParallelSum: %d", seqSum(), totalSum)
}

func seqSum() int {
	var sum int
	for i := 0; i < N; i++ {
		sum += numArray[i]
	}
	return sum
}
