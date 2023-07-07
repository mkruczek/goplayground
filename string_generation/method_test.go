package string_generation

import "testing"

//Benchmark_Bytes-8           	 4370394	       265.9 ns/op
//Benchmark_Runes-8           	 3285606	       352.8 ns/op
//Benchmark_StringBuilder-8   	 4270227	       277.1 ns/op

func Benchmark_Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteGenerateString(10)
	}
}

func Benchmark_Runes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runeGenerateString(10)
	}
}

func Benchmark_StringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuilderGenerateString(10)
	}
}
