package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println("Error Dialing", err.Error())
	}
	defer conn.Close()

	_, err = conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	if err != nil {
		fmt.Println("Error Writing", err.Error())

		return
	}

	buff := make([]byte, 1024)

	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error Reading", err.Error())
	}
	fmt.Println(string(buff[:n]))
}
