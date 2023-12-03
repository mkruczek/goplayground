package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)

func main() {

	go func() {
		for {
			x, y := robotgo.Location()
			robotgo.Move(x+1, y+1)
			st := sleepTime()
			time.Sleep(st)
		}
	}()

	done := make(chan bool)
	go func() {
		for {
			x, y := robotgo.Location()
			if x == 0 && y == 1079 {
				fmt.Println("done")
				done <- true
			}
		}
	}()

	<-done
}

// sleepTime - return random value between 240 and 300
func sleepTime() time.Duration {
	st := 240 + rand.Intn(60)
	return time.Duration(st) * time.Second
}
