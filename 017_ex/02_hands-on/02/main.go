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
		}

		defer conn.Close()


			//we not go here because scanner scan over and over open stream
			//no exit instruction
			//scan always in this example before close
			fmt.Println("Code got here.")
			io.WriteString(conn, "I see you connected.")

		conn.Close()

		}
}


