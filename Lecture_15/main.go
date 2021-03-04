package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("Net_server is running")

	nl, err := net.Listen("tcp", "localhost:8888")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	conn, err := nl.Accept()

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(1)
	}

	//to take remote address

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(remoteAddr)

	bs := make([]byte, 1024)

	conn.Read(bs)
	fmt.Println(string(bs))

	conn.Write([]byte("Welcome to Auhidul Islam Siam."))

	conn.Close()

	nl.Close()

}
