package main

import (
	"fmt"
	"unsafe"
)

type singleMethod interface {
	first()
}

type tenMethod interface {
	first()
	second()
	third()
	fourth()
	fifth()
	sixth()
	seventh()
	eighth()
	ninth()
	tenth()
}

type object struct{}

func (o object) first()   {}
func (o object) second()  {}
func (o object) third()   {}
func (o object) fourth()  {}
func (o object) fifth()   {}
func (o object) sixth()   {}
func (o object) seventh() {}
func (o object) eighth()  {}
func (o object) ninth()   {}
func (o object) tenth()   {}

func main() {

	var sm singleMethod
	sm = object{}

	var tm tenMethod
	tm = object{}

	fmt.Printf("single: type %T with size %d\n", sm, unsafe.Sizeof(sm))
	fmt.Printf("ten: type %T with size %d\n", tm, unsafe.Sizeof(tm))

}
