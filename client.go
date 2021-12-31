package main

import (
	"bufio"
	"fmt"
	"github.com/k1le0/demo/internal"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	//
	conn, err := net.Dial("tcp", ":8090")
	if err != nil {
		fmt.Printf("Connect error!")
		return
	}

	defer conn.Close()

	sendMsg(conn)
}

func sendMsg(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read Input Error, err: ", err)
		}
		content := strings.TrimSpace(input)
		message := internal.Message{
			UserName: "client1",
			Content:  content,
			Time:     "123",
		}

		_, err = conn.Write(internal.EncodeToBytes(message))
		if err != nil {
			fmt.Printf("Write failed, err: %s", err)
			break
		}
		time.Sleep(5 * time.Second)
	}
}
