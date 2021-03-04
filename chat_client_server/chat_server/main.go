package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	nl, err := net.Listen("tcp", "localhost:8888")

	if err != nil {

		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer nl.Close()
	log.Printf("Server started at localhost:8888")
	for {

		conn, err := nl.Accept()

		if err != nil {

			fmt.Println(err.Error())
			os.Exit(1)
		}

		bs := make([]byte, 1024)

		n, err := conn.Read(bs)

		if err != nil {

			fmt.Println(err.Error())
			//continue
		}

		rstring := string(bs[:n])
		fmt.Println(rstring)

		rectime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(rectime)

		msg := fmt.Sprintf(`Your massage is %s and received at %s`, rstring, rectime)
		fmt.Println(msg)

		conn.Write([]byte(msg))

	}

}
