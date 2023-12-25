package christmastree

import (
	"fmt"
	"strings"
)

func printTree(levels int) {
	for i := 0; i < levels; i++ {
		fmt.Print(strings.Repeat(" ", levels-i-1))
		fmt.Println(strings.Repeat("*", 2*i+1))
	}

	// print the trunk of the tree //todo add levels and width base on the tree size
	fmt.Print(strings.Repeat(" ", levels-1))
	fmt.Println("*")
}
