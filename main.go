package main

import (
	"fmt"

	"github.com/alvinend/proto-go/src/jsonutils"
	"github.com/alvinend/proto-go/src/reader"
	"github.com/alvinend/proto-go/src/simplepb"
	"github.com/alvinend/proto-go/src/writer"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

func main() {
	sm := doSimple()

	readAndWrite(sm)
	jsonReadAndWrite(sm)
}

// JSON Demo
func jsonReadAndWrite(sm protoiface.MessageV1) {
	smAsString := jsonutils.ToJSON(sm)
	fmt.Println(smAsString)
	sm2 := &simplepb.SimpleMessage{}
	jsonutils.FromJSON(smAsString, sm2)
	fmt.Println("From JSON", sm2)
}

// Read and Write Demo
func readAndWrite(sm proto.Message) {
	writer.WriteToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	reader.ReadFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

// Simple Thing To Do
func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 2, 3},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm)

	fmt.Println("The ID is", sm.GetId())

	return &sm
}
