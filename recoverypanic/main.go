package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for {
			fmt.Println("goroutine")
			time.Sleep(1 * time.Second)
		}
	}()

	panicWithRecovery()

	done := make(chan bool)
	<-done

}

// func will panic
func panicWithRecovery() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic")
		}
	}()

	callPanic()
}

func callPanic() {
	panic("panic")
}

//func setSomeValue
