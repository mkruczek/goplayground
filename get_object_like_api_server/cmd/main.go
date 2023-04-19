package main

import (
	get_object_like_api_server "getobject"
	"os"
	"runtime/trace"
)

func main() {

	f, err := os.Create("trace_value.out")
	//f, err := os.Create("trace_pointer.out")
	if err != nil {
		panic(err)
	}

	trace.Start(f)
	defer trace.Stop()

	n := 1_000_000
	get_object_like_api_server.MainFunctionByValue(n)
	//get_object_like_api_server.MainFunctionByPointer(n)
}
