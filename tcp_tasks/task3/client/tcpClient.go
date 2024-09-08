package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)	// This establishes a TCP connection
	if err != nil {
		fmt.Println("err")
		return
	}
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter action (reverse/removeSpaces) || message: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if strings.HasSuffix(text, "STOP") {
			fmt.Println("TCP client exiting...")
			return
		}

		fmt.Fprintf(c, text + "\n")

		message, _ := bufio.NewReader(c).ReadString('\n')	// Reading data from a TCP connection
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		fmt.Printf("-> Server response: %s", message)
	}
}