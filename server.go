package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("starting the server ...")
	// create listener
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	//listen connection from client
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Receving data: %v", string(buf[:len]))
	}
}
