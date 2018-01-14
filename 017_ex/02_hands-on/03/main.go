package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
			if line == "" {
				fmt.Println("Empty line")
				break
			}
		}

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")

		conn.Close()

	}
}


