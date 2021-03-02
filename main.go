package main

import (
	"fmt"

	simplepb "github.com/stephenjlovell/protobuf-example-go/src/simple"
)

func main() {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "my simple message",
		SampleList: []int32{1, 1, 2, 3, 5},
	}
	fmt.Printf("%+v\n", sm)
	fmt.Printf(sm.GetId())

}
