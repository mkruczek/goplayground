package bench

import "testing"

func Benchmark_Pointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newPointer := NewPointer()

		newPointer.Add(1, 2)
		newPointer.Sub(1, 2)
		newPointer.Mul(1, 2)
		newPointer.Div(1, 2)
	}
}

func Benchmark_Value(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newValue := NewValue()

		newValue.Add(1, 2)
		newValue.Sub(1, 2)
		newValue.Mul(1, 2)
		newValue.Div(1, 2)
	}
}

type Value struct{}

func NewValue() Value { return Value{} }

func (v Value) Add(a int, b int) int { return a + b }

func (v Value) Sub(a int, b int) int { return a - b }

func (v Value) Mul(a int, b int) int { return a * b }

func (v Value) Div(a int, b int) int { return a / b }

type Pointer struct{}

func NewPointer() *Pointer { return &Pointer{} }

func (p *Pointer) Add(a int, b int) int { return a + b }

func (p *Pointer) Sub(a int, b int) int { return a - b }

func (p *Pointer) Mul(a int, b int) int { return a * b }

func (p *Pointer) Div(a int, b int) int { return a / b }
