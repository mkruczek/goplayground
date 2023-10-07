package algoSerach

import (
	"sync"
	"testing"
)

/*
TLDR: Parallel search may be faster than linear search, but it depends on the case, and the difference is not significant, so it's not worth it.

Results:
case: target = 999_999 (bad case for linear search); numWorkers = 10, len(slice) = 1_000_000
Benchmark_LinearSearch-8           	     518	   2083226 ns/op	 8003593 B/op	       1 allocs/op
Benchmark_ParallelLinearSearch-8   	     660	   1742301 ns/op	 8006433 B/op	      21 allocs/op

case: target = 2 (good case for linear search); numWorkers = 100, len(slice) = 1_000_000
Benchmark_LinearSearch-8           	     886	   1297899 ns/op	 8003592 B/op	       1 allocs/op
Benchmark_ParallelLinearSearch-8   	     723	   2507836 ns/op	 8009922 B/op	      33 allocs/op

case: target = 50 (average case for linear search); numWorkers = 10, len(slice) = 100
Benchmark_LinearSearch-8           	14567437	        68.99 ns/op	       0 B/op	       0 allocs/op
Benchmark_ParallelLinearSearch-8   	   71700	     18457 ns/op	    4643 B/op	      24 allocs/op
*/

const target = 50

func Benchmark_LinearSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linearSearch(returnSliceWith100Elements(), target)
	}
}

func Benchmark_ParallelLinearSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallelSearch(returnSliceWith100Elements(), target, 10)
	}
}

// linearSearch O(n)
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// parallelSearch O(n) + context switch overhead :D
func parallelSearch(arr []int, target int, numWorkers int) int {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	subarraySize := len(arr) / numWorkers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startIdx := i * subarraySize
		endIdx := (i + 1) * subarraySize

		go searchInSubarray(arr, target, startIdx, endIdx, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		if result != -1 {
			return result
		}
	}

	return -1
}

func searchInSubarray(arr []int, target int, startIdx int, endIdx int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for i := startIdx; i < endIdx; i++ {
		if arr[i] == target {
			ch <- arr[i]
			return
		}
	}
	ch <- -1
}

func returnSliceWith1_000_000Elements() []int {
	s := make([]int, 1_000_000)
	for i := range s {
		s[i] = i
	}
	return s
}

func returnSliceWith100Elements() []int {
	s := make([]int, 100)
	for i := range s {
		s[i] = i
	}
	return s
}
