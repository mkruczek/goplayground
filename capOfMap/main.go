package main

import (
	"fmt"
	"unsafe"
)

//todo -> work in progress
// https://github.com/golang/go/blob/cd99385ff4a4b7534c71bb92420da6f462c5598e/src/runtime/map.go#L115
// https://levelup.gitconnected.com/memory-allocation-and-performance-in-golang-maps-b267b5ad9217
// https://hackernoon.com/some-insights-on-maps-in-golang-rm5v3ywh

func main() {

	m := make(map[int]struct{})
	ptr := &m
	fmt.Printf("map address: %p\n", ptr)

	for i := 0; i < 1000; i++ {
		m[i] = struct{}{}
		size := int(unsafe.Sizeof(m))
		fmt.Println("size: ", size) //not even close :facepalm:
	}

	ptr = &m
	fmt.Printf("map address: %p\n", ptr)
}
