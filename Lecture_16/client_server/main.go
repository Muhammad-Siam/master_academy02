package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")

	if err != nil {
		fmt.Println(err.Error())
	}
	conn.Write([]byte("Hello, This is Siam."))

	bs := make([]byte, 1024)

	n, _ := conn.Read(bs)

	fmt.Println(string(bs[:n]))

	conn.Close()

}
