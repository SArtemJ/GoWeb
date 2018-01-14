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
			continue
		}

		callfunc(conn)
	}
}


func callfunc(conn net.Conn) {
	defer conn.Close()
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
			if line == "" {
				fmt.Println("Empty line")
				break
			}
		}
		body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
		//status line
		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		//response headers
		fmt.Fprintf(conn, "Content-lenght: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-type: text/plain\r\n")
		io.WriteString(conn, "\r\n")
		io.WriteString(conn, body)
}

