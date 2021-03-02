package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	simplepb "github.com/stephenjlovell/protobuf-example-go/src/simple"
)

func main() {
	readWriteDemo()
}

func readWriteDemo() {
	fname := "simple.bin"
	sm := doSimple()
	writeToFile(fname, sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile(fname, sm2)
	fmt.Println("read data from disk:", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	bytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("failed to serialize message", err)
		return err
	}

	if err := ioutil.WriteFile(fname, bytes, 0644); err != nil {
		log.Fatalln("failed to write to disk", err)
		return err
	}
	fmt.Println("data written to disk")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("failed to read from file", err)
		return err
	}
	if err := proto.Unmarshal(bytes, pb); err != nil {
		log.Fatalln("failed to deserialize message", err)
		return err
	}
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "my simple message",
		SampleList: []int32{1, 1, 2, 3, 5},
	}
	fmt.Printf("%+v\n", sm)
	return &sm
}
