package main

import "fmt"

func main() {

	var x uint8 = 32

	fmt.Printf("x=%d // decimal\n", x)
	fmt.Printf("x=%b // binary \n", x)
	fmt.Printf("x=%o // octal \n", x)
	fmt.Printf("x=%x // hexadecimal \n", x)
	fmt.Printf("x=%X // hexadecimal \n", x)

	fmt.Printf("x=%c // char \n", x)
	fmt.Printf("x=%q // quoted char \n", x)
}
