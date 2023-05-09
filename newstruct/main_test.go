package newstruct

import "testing"

type S struct {
	A int
	B string
}

func (s S) doNothing() {

}

func NewS_Value() S {
	return S{
		A: 0,
		B: "0",
	}
}

func NewS_ValueWithPreDeclaration() S {
	s := S{
		A: 0,
		B: "0",
	}
	return s
}

func NewS_Pointer() *S {
	return &S{
		A: 0,
		B: "0",
	}
}

func NewS_PointerWithPreDeclaration() *S {
	s := S{
		A: 0,
		B: "0",
	}

	return &s
}

func Benchmark_Value(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < b.N; i++ {
			s := NewS_Value()
			s.doNothing()
		}
	}
}

func Benchmark_ValueWithPreDeclaration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < b.N; i++ {
			s := NewS_ValueWithPreDeclaration()
			s.doNothing()
		}
	}
}

func Benchmark_Pointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < b.N; i++ {
			s := NewS_Pointer()
			s.doNothing()
		}
	}
}

func Benchmark_PointerWithPreDeclaration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < b.N; i++ {
			s := NewS_PointerWithPreDeclaration()
			s.doNothing()
		}
	}
}
