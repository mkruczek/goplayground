package main

import (
	"fmt"
	"unsafe"
)

type doc struct {
	id    int
	owner string
	value []byte
}

func main() {

	//doc with byte array of 1MB
	doc1 := doc{1, "owner1", make([]byte, 1024*1024)}

	//doc with byte array of 2MB
	doc2 := doc{2, "owner2", make([]byte, 1024*1024*2)}

	//doc with byte array of 3MB
	doc3 := doc{3, "owner3", make([]byte, 1024*1024*3)}

	//doc with byte array of 4MB
	doc4 := doc{4, "owner4", make([]byte, 1024*1024*4)}

	fmt.Println("Size of doc1: ", unsafe.Sizeof(doc1))
	fmt.Println("Size of doc2: ", unsafe.Sizeof(doc2))
	fmt.Println("Size of doc3: ", unsafe.Sizeof(doc3))
	fmt.Println("Size of doc4: ", unsafe.Sizeof(doc4))

	//size of doc1 value
	fmt.Println("Size of doc1 value: ", unsafe.Sizeof(doc1.value))

	//size of doc2 value
	fmt.Println("Size of doc2 value: ", unsafe.Sizeof(doc2.value))

	//size of doc3 value
	fmt.Println("Size of doc3 value: ", unsafe.Sizeof(doc3.value))

	//size of doc4 value
	fmt.Println("Size of doc4 value: ", unsafe.Sizeof(doc4.value))
}
