package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)	// TCP listens for incoming connections
	if err != nil {
		fmt.Println("err")
		return
	}
	defer l.Close()

	c, err := l.Accept()	// Server accepts a connection from a client
	if err != nil {
		fmt.Println("err")
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("err")
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("->: " + string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		// msg := "Hello Mark"
		// fmt.Println(msg)
		c.Write([]byte(myTime))	  // Writing data back to the client

	}
}