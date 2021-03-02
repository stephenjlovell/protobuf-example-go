package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	simplepb "github.com/stephenjlovell/protobuf-example-go/src/simple"
)

func main() {
	sm := makeSimple()
	readWriteDemo(sm)
	serializeDemo(sm)
	enumDemo()
}

func enumDemo() {
	// em := &enumpb.EnumMessage{}
}

func serializeDemo(pb proto.Message) {
	str, _ := toJSON(pb)
	fmt.Println(str)
	sm2 := &simplepb.SimpleMessage{}
	fromJSON(str, sm2)
	fmt.Println("successfully created object:", sm2)
}

func toJSON(pb proto.Message) (string, error) {
	m := jsonpb.Marshaler{}
	str, err := m.MarshalToString(pb)
	if err != nil {
		log.Fatalln("could not serialize object to JSON", err)
		return "", err
	}
	return str, nil
}

func fromJSON(str string, pb proto.Message) error {
	if err := jsonpb.UnmarshalString(str, pb); err != nil {
		log.Fatalln("could not unserialize object from JSON", err)
		return err
	}
	return nil
}

func readWriteDemo(sm proto.Message) {
	fname := "simple.bin"
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

func makeSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "my simple message",
		SampleList: []int32{1, 1, 2, 3, 5},
	}
	fmt.Printf("%+v\n", sm)
	return &sm
}
