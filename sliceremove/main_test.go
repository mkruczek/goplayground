package main

import "testing"

func orderingRemove(s []string, i int) []string {
	result := append(s[:i], s[i+1:]...)
	return result
}

func shiftRemove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	result := s[:len(s)-1]
	return result
}

/*
goos: linux
goarch: amd64
pkg: some-benchmark/sliceremove
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkShiftRemove
BenchmarkShiftRemove-8   	172545070	         6.884 ns/op
BenchmarkPop
BenchmarkPop-8           	1000000000	         0.6414 ns/op
*/

func BenchmarkOrderRemove(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for i := 0; i < b.N; i++ {
		orderingRemove(s, 2)
	}
}

func BenchmarkShiftRemove(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for i := 0; i < b.N; i++ {
		shiftRemove(s, 2)
	}
}
