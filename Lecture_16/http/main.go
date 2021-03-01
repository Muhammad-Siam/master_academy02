package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("Web server")

	nl, err := net.Listen("tcp", "localhost:8888")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

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

		fmt.Println(n)

		reqstr := string(bs)
		fmt.Println(reqstr)

		//body

		body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Siam</title>
	</head>
	<body>
		<h1 align= center>Myself</h1>
		<p>Hello, This is Siam</p>
	</body>
	</html>`

		/*fmt.Fprint(conn, "HTTP/ 1.1 200 OK\r\n")
		fmt.Fprintf(conn, "content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "content-type: text/html \r \n")
		fmt.Fprint(conn, "\r\n")*/
		resp := "HTTP/1.1 200 OK\r\n" +
			"content-Length: %d\r\n" +
			"content-Type: text/html \r\n" +
			"\r\n%s"

		msg := fmt.Sprintf(resp, len(body), body)

		fmt.Println(msg)
		conn.Write([]byte(msg))
	}

	//conn.Close()

}
