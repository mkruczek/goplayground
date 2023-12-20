package main

import (
	"fmt"
)

func main() {

	if err := tryCatch(myPanic, "test"); err != nil {
		fmt.Println(err)
	}
}

func tryCatch(args ...any) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("revovery: %s", e)
		}
	}()

	for _, arg := range args {
		switch arg.(type) {
		case func():
			arg.(func())()
		case string:
			fmt.Println(arg)
		}
	}

	return nil
}

func myPanic() {
	panic("my panic")
}
