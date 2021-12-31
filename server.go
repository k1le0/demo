package main

import (
	"fmt"
	"github.com/k1le0/demo/internal"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Printf("tcp监听8090失败")
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	defer conn.Close()
	for {
		var buf = make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read failed, err: ", err)
			break
		}
		data := internal.DecodeToMessage(buf[:n])
		fmt.Println("Read from ", data.UserName)
		fmt.Println("Read Content ", data.Content)
		fmt.Println("----------------------------")
	}
}
