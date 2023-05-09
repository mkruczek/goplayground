package main

//go tool compile -m ./main_test.go

func main() {
	obj := NewPointer()
	//obj := NewValue()

	obj.Add(1, 2)
	obj.Sub(1, 2)
	obj.Mul(1, 2)
	obj.Div(1, 2)
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
