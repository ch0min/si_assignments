package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// Function to reverse a string
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Function to remove all white space from a string
func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

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

	for {
		c, err := l.Accept()	// Server accepts a connection from a client
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading data:", err)
			return
		}
		netData = strings.TrimSpace(netData)
		fmt.Printf("Received: %s\n", netData)

		if netData == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		// Split the message into action and text
		parts := strings.SplitN(netData, "||", 2)
		if len(parts) != 2 {
			c.Write([]byte("Invalid format. Please use action||message\n"))
			continue
		}
		action, message := parts[0], parts[1]

		// Determine the action and process the message
		var result string
		switch action {
		case "reverse":
			result = reverse(message)
		case "removeSpaces":
			result = removeSpaces(message)
		default:
			result = "Unknown action. Use 'reverse' or 'removeSpaces'."
		}

		// Send the result back to the client
		c.Write([]byte(result + "\n"))
		t := time.Now()
		fmt.Printf("Processed at: %s\n", t.Format(time.RFC3339))
	}
}