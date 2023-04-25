package main

import "testing"

//BenchmarkDuck-8 [interface bird]   	    7455	    162093 ns/op	       0 B/op	       0 allocs/op
//BenchmarkDuck-8 [struct duck]  	       16732	     62525 ns/op	       0 B/op	       0 allocs/op

func BenchmarkDuck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		undertest()
	}
}
