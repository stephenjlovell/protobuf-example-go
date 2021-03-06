package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	tutorial "github.com/stephenjlovell/protobuf-example-go/src/addressbook"
	complexpb "github.com/stephenjlovell/protobuf-example-go/src/complex"
	enumpb "github.com/stephenjlovell/protobuf-example-go/src/enum_example"
	simplepb "github.com/stephenjlovell/protobuf-example-go/src/simple"
)

func main() {
	sm := makeSimple()
	readWriteDemo(sm)
	serializeDemo(sm)
	enumDemo()
	complexDemo()
	personDemo()
}

func personDemo() {
	pb := tutorial.AddressBook{
		People: []*tutorial.Person{
			makePerson(1, "Willow"),
			makePerson(1, "Xander"),
			makePerson(1, "Tara"),
		},
	}
	fmt.Println(pb)
}

func makePerson(id int32, name string) *tutorial.Person {
	return &tutorial.Person{
		Id:    id,
		Name:  name,
		Email: fmt.Sprintf("%s@example.com", name),
		Phones: []*tutorial.Person_PhoneNumber{
			{
				Number: "111-222-3333",
				Type:   tutorial.Person_MOBILE,
			},
			{
				Number: "111-222-5555",
				Type:   tutorial.Person_WORK,
			},
		},
	}
}

func complexDemo() {
	com := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "main message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   2,
				Name: "message 2",
			},
			{
				Id:   3,
				Name: "message 3",
			},
			{
				Id:   4,
				Name: "message 4",
			},
		},
	}

	fmt.Println(com)
}

func enumDemo() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_FRIDAY,
	}
	fmt.Println(em)
	fmt.Printf("yay, it's %s\n", em.GetDayOfTheWeek().String())
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
