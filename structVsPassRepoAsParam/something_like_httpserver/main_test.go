package something_like_httpserver

import "testing"

//goos: windows
//goarch: amd64
//pkg: temporary
//cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
//Benchmark_Struct-12      4552531               270.8 ns/op            32 B/op 2 allocs/op
//Benchmark_Func-12        5361493               235.6 ns/op            24 B/op 1 allocs/op

func Benchmark_Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainStruct()
	}
}

func Benchmark_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainFunc()
	}
}
