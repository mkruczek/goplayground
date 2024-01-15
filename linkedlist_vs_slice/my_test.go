package linkedlist_vs_slice

import "testing"

//benchmark tests for linked list vs array

func Benchmark_LinkedList(b *testing.B) {

	linked := LinkedList{}

	for i := 0; i < b.N; i++ {
		linked.Add(i)
		linked.Remove(i)

	}
}

func Benchmark_Slice(b *testing.B) {

	slice := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		slice[i] = i
		slice[i] = 0

	}
}
