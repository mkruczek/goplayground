package main

type myInter interface {
	doSomething()
}

type myStruct struct {
}

func (m myStruct) doSomething() {
}

func forMyInter(mi myInter) {
	mi.doSomething()
}

func forMyStruct(ms myStruct) {
	ms.doSomething()
}

func main() {
	forMyInter(myStruct{})
	//forMyStruct(myStruct{})
}
