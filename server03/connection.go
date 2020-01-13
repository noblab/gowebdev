package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}
	fmt.Println("Code got here.")
	io.WriteString(conn, "I see you connected.")
}
