package switchvsmap

import "testing"

// example comes from exercises/resistor-color

var resistanceMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func ColorCodeMap(color string) int {
	return resistanceMap[color]
}

func ColorCodeSwitch(color string) int {
	switch color {
	case "black":
		return 0
	case "brown":
		return 1
	case "red":
		return 2
	case "orange":
		return 3
	case "yellow":
		return 4
	case "green":
		return 5
	case "blue":
		return 6
	case "violet":
		return 7
	case "grey":
		return 8
	case "white":
		return 9
	default:
		return -1
	}
}

//Benchmark_Map-8      	24996829	        41.94 ns/op	       0 B/op	       0 allocs/op
//Benchmark_Switch-8   	252630018	         4.667 ns/op	       0 B/op	       0 allocs/op

func Benchmark_Map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ColorCodeMap("red")
		ColorCodeMap("blue")
		ColorCodeMap("green")
		ColorCodeMap("yellow")
	}
}

func Benchmark_Switch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ColorCodeSwitch("red")
		ColorCodeSwitch("blue")
		ColorCodeSwitch("green")
		ColorCodeSwitch("yellow")
	}
}

// some other stuff like a curiosity: slice can be used as a map but key must be int
var sliceCode = []string{
	0: "black",
	1: "brown",
	2: "red",
	3: "orange",
	4: "yellow",
	5: "green",
	6: "blue",
	7: "violet",
	8: "grey",
	9: "white",
}

func ColorCodeSlice(value int) string {
	return sliceCode[value]
}

// Benchmark_Slice-8    	1000000000	         0.9685 ns/op	       0 B/op	       0 allocs/op
func Benchmark_Slice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ColorCodeSlice(0)
		ColorCodeSlice(3)
		ColorCodeSlice(5)
		ColorCodeSlice(9)
	}
}
