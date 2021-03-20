package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	resp := []byte("\x00" + "rootersiam2255@gmail.com" + "\x00" + "AuHIDS19@gmail.com")
	fmt.Println(string(resp), resp)

	sEnc := base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println(sEnc)

}
