package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	serverAddr = "127.0.0.1:8080"
)

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	go func() {
		for {
			msg := make([]byte, 4096)
			n, err := conn.Read(msg)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(strings.TrimSpace(string(msg[:n])))
		}
	}()

	for {
		var msg string
		fmt.Scanln(&msg)
		if msg == "/quit" {
			return
		}

		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
