package jsonutils

import (
	"log"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/runtime/protoiface"
)

func ToJSON(pb protoiface.MessageV1) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Cant convert to JSON", err)
		return ""
	}
	return out
}

func FromJSON(in string, pb protoiface.MessageV1) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Cant Convert shit", err)
	}
}
