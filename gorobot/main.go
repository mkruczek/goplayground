package main

import (
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	for {
		x, y := robotgo.Location()
		robotgo.Move(x+1, y+1)
		time.Sleep(60 * time.Second)
	}
}
