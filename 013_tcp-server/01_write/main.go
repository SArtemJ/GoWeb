package main

import (
	"net"
	"log"
	"io"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintf(conn, "How is your day\n")
		fmt.Fprintf(conn, "%v", "Well, I hope\n")

		conn.Close()
	}
}
