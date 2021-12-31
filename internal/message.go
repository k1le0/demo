package internal

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Message struct {
	UserName string
	Content  string
	Time     string
}

func EncodeToBytes(message Message) []byte {
	buf := bytes.Buffer{}
	encode := gob.NewEncoder(&buf)
	err := encode.Encode(message)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func DecodeToMessage(b []byte) Message {
	message := Message{}
	decode := gob.NewDecoder(bytes.NewReader(b))
	err := decode.Decode(&message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}
