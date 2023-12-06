// JD Waldron and Braydon Hughes
// Program that sums up an array concurrently
package main

import (
	"fmt"
	"sync"
)

var (
	numArray [10000]int     // array to sum up
	totalSum int            /// parallel sum of the array
	mu       sync.Mutex     // mutex tool
	wg       sync.WaitGroup // helps organize all goroutines
	N        int            // items in the array
)

func work(low int, high int) {
	var localSum int // sum for this portion of the array
	// each portion does its own work
	for i := low; i < high; i++ {
		localSum += numArray[i]
	}
	mu.Lock()            // lock totalSum
	totalSum += localSum // goroutine updates the global
	mu.Unlock()          // unlock totalSum
	wg.Done()            // mark that this current goroutine is done with its work
}

func main() {
	numThreads := 10  // numThreads
	totalSum = 0      // set totalSum
	N = len(numArray) // set N

	// initialize array values
	for i := 0; i < N; i++ {
		numArray[i] = i + 1
	}

	// sections in the array
	workNum := N / numThreads

	// assigns different goroutines to their sections
	for i := 0; i < numThreads; i++ {
		low := i * workNum        // computes starting point of each goroutine
		high := (i + 1) * workNum // computes stopping point of each goroutine
		wg.Add(1)                 // mark that a gorountine has been started
		go work(low, high)        // calls helper function for each goroutine
	}

	wg.Wait() // wait for all goroutines to finish

	fmt.Printf("SeqSum: %d, ParallelSum: %d", seqSum(), totalSum) // print sequential and parallel sums
}

// calculate the sum of the array sequentially
func seqSum() int {
	var sum int
	for i := 0; i < N; i++ {
		sum += numArray[i]
	}
	return sum
}
