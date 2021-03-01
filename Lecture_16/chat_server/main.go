package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	fmt.Println("This is a chat server")

	nl, err := net.Listen("tcp", "localhost:8888")

	if err != nil {

		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer nl.Close()
	log.Printf("Server started on localhost:8888")

	for {

		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
			//continue
		}
		bs := make([]byte, 1024)

		n, err := conn.Read(bs)

		if err != nil {
			fmt.Println(err.Error())

		}
		rstr := string(bs[:n])
		fmt.Println(rstr)

		rectime := time.Now().Format("2006-01-02 15:04:05")

		msg := fmt.Sprintf(`your massage is %s, received at %s`, rstr, rectime)

		fmt.Println(msg)
		conn.Write([]byte(msg))
	}
}
