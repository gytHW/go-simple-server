package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	server = "127.0.0.1:5000"
)

func main() {
	//open connection
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what's your name?")
	clientName, _ := inputReader.ReadString('\n')
	fmt.Printf("clientName is %s", clientName)
	trimmedClient := strings.Trim(clientName, "\r\n") //windows, if linux use "\n" instead

	for {
		fmt.Println("what to send to the server? Type q to quit")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "q" {
			return
		}

		_, err := conn.Write([]byte(trimmedClient + " says:" + trimmedInput + "\n"))
		if err != nil {
			fmt.Println("Error writing", err.Error())
			return
		}
	}
}
