package caller

import "testing"

/*
goos: linux
goarch: amd64
pkg: some-benchmark/structVsPassRepoAsParam/caller
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkServiceStruct
BenchmarkServiceStruct-8            	  218053	      5641 ns/op	       0 B/op	       0 allocs/op
BenchmarkServicePassRepoAsParam
BenchmarkServicePassRepoAsParam-8   	  254794	      4681 ns/op	       0 B/op	       0 allocs/op
*/

func BenchmarkServiceStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callerStructWithRepo()
	}
}

func BenchmarkServicePassRepoAsParam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callerPassRepoAsParam()
	}
}
