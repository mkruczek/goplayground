package main

import (
	"testing"
)

// https://ueokande.github.io/go-slice-tricks/
func orderingRemove(s []string, i int) []string {
	result := append(s[:i], s[i+1:]...)
	return result
}

func shiftRemove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	result := s[:len(s)-1]
	return result
}

func orderWithCopy(s []string, i int) []string {
	result := s[:i+copy(s[i:], s[i+1:])]
	return result
}

/*
goos: linux
goarch: amd64
pkg: some-benchmark/sliceremove
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkOrderRemove
BenchmarkOrderRemove-8   	160714988	         7.300 ns/op	       0 B/op	       0 allocs/op
BenchmarkShiftRemove
BenchmarkShiftRemove-8   	1000000000	         0.6052 ns/op	       0 B/op	       0 allocs/op
BenchmarkTestName
BenchmarkTestName-8      	165413005	         7.128 ns/op	       0 B/op	       0 allocs/op
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

func BenchmarkTestName(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for i := 0; i < b.N; i++ {
		orderWithCopy(s, 2)
	}
}
