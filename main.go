package main

import (
	"fmt"
	"github.com/k1le0/demo/internal"
)

func main() {
	message := internal.Message{
		UserName: "client",
		Content:  "123",
		Time:     "123",
	}
	b := internal.EncodeToBytes(message)
	fmt.Println("123: ", b)
	d := internal.DecodeToMessage(b)
	fmt.Println("456: ", d.Content)
}
