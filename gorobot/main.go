package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {

	go func() {
		for {
			x, y := robotgo.Location()
			robotgo.Move(x+1, y+1)
			time.Sleep(60 * time.Second)
		}
	}()

	done := make(chan bool)
	go func() {
		for {
			x, y := robotgo.Location()
			//fmt.Println("x:", x, "y:", y)
			//time.Sleep(1 * time.Second)
			if x == 0 && y == 1079 {
				fmt.Println("done")
				done <- true
			}
		}
	}()

	<-done

}
