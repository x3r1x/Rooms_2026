package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	serverAddr := "84.201.159.214:8080"
	connection, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("You write message to server: ", text)
		if strings.ToLower(text) == "exit" {
			break
		}

		_, err := connection.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, err := connection.Read(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("The golang Gopher answer you: ", string(buffer[:n]))
	}
}
