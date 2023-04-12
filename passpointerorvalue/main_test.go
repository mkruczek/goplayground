package passpointerorvalue

import "testing"

type T struct {
	i int
	s string
}

func (t *T) SetI(i int) {
	t.i = i
}

func (t *T) SetS(s string) {
	t.s = s
}

func (t T) SetIReturnT(i int) T {
	t.i = i
	return t
}

func (t T) SetSReturnT(s string) T {
	t.s = s
	return t
}

func Benchmark_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := T{}
		t.SetI(1)
		t.SetS("a")
	}
}

func Benchmark_SetAtPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := &T{}
		t.SetI(1)
		t.SetS("a")
	}
}

func Benchmark_SetReturn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := T{}
		t = t.SetIReturnT(1)
		t = t.SetSReturnT("a")
	}
}
