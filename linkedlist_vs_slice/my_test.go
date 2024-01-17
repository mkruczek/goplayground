package linkedlist_vs_slice

import (
	"testing"
)

//for size = 100_000
//Benchmark_Array-8        	1000000000	         0.01323 ns/op
//Benchmark_Slice-8        	1000000000	         0.01188 ns/op
//Benchmark_LinkedList-8   	1000000000	         0.02830 ns/op

const size = 100_000

// benchmark tests for linked list vs array
func Benchmark_Array(b *testing.B) {

	var array [size]int

	for i := 0; i < size; i++ {
		array[i] = i
	}

	for i := 0; i < size; i++ {
		array[i] = i
		array[i] = 0

		if i%1000 == 0 {
			//remove element with index i / 2; replace with 0
			array[i/2] = 0
			//search for element with value i / 3
			for j := 0; j < len(array); j++ {
				if array[j] == i/3 {
					break
				}
			}
		}
	}
}

func Benchmark_Slice(b *testing.B) {

	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = i
	}

	for i := 0; i < len(slice); i++ {
		slice[i] = i
		slice[i] = 0

		if i%1000 == 0 {
			//remove element with index i / 2
			slice = append(slice[:i/2], slice[i/2+1:]...)
			//search for element with value i / 3
			for j := 0; j < len(slice); j++ {
				if slice[j] == i/3 {
					break
				}
			}
		}
	}
}

func Benchmark_LinkedList(b *testing.B) {
	list := LinkedList{}

	for i := 0; i < size; i++ {
		list.Add(i)
	}

	for i := 0; i < size; i++ {
		list.Add(i)
		list.Remove(i)

		if i%1000 == 0 {
			//remove element with index i / 2
			list.Remove(i / 2)
			//search for element with value i / 3
			list.Search(i / 3)
		}

	}

}
