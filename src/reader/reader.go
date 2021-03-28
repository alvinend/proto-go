package reader

import (
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func ReadFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Cant read shit")
		return err
	}

	err2 := proto.Unmarshal(in, pb)

	if err2 != nil {
		log.Fatalln("Cant understand shit")
		return err2
	}

	return nil
}
