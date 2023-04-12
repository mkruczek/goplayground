package get_object_from_repository

import "testing"

/*
goos: linux
goarch: 386
pkg: some-benchmark/get_object_from_repository
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
Benchmark_Value
Benchmark_Value-8     	 7952400	       140.7 ns/op	      32 B/op	       2 allocs/op
Benchmark_Pointer
Benchmark_Pointer-8   	 8608419	       135.7 ns/op	      32 B/op	       2 allocs/op
*/

func Benchmark_Value(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainFunctionByValue()
	}
}

func Benchmark_Pointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainFunctionByPointer()
	}
}
