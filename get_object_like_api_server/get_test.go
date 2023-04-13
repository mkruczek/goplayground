package get_object_like_api_server

import "testing"

/*
goos: linux
goarch: 386
pkg: some-benchmark/get_object_like_api_server
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
Benchmark_Value
Benchmark_Value-8     	    6880	    178440 ns/op	       9 B/op	       0 allocs/op
Benchmark_Pointer
Benchmark_Pointer-8   	    6693	    181354 ns/op	       9 B/op	       0 allocs/op
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
