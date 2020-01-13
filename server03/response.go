package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
		defer conn.Close()

		request(conn)
	}
}

func respondIndex(conn net.Conn) {
	body := `<!DOCTYPE html><header lang ="ja"><meta charset = "UTF-8"><title>Response</title></header><body><a href = "localhost:8080"><h1>Index<h1></a><a href="/apply"><h1>Apply</h1></a></body>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func respondApply(conn net.Conn) {
	body := `<!DOCTYPE html><header lang ="ja"><meta charset = "UTF-8"><title>Response</title></header><body><a href = "localhost:8080"><h1>Index<h1></a><a href="localhost:8080/apply"><h1>Apply</h1></a><form action ="/apply" method="POST"><input type="submit" value="submit"></form></body></html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func respondPost(conn net.Conn) {
	body := `<!DOCTYPE html><header lang ="ja"><meta charset = "UTF-8"><title>Response</title></header><body><h1>POST APPLY</h1><a href = "localhost:8080"><h1>Index<h1></a><a href="localhost:8080/apply"><h1>Apply</h1></a></body>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func request(conn net.Conn) {
	sc := bufio.NewScanner(conn)
	count := 0
	for sc.Scan() {
		ln := sc.Text()
		fs := strings.Fields(ln)
		if count == 0 {
			//fmt.Printf("METHOD***%v\n", fs[0])
			switch fs[0] {
			case "GET":
				if fs[1] == "/apply" {
					respondApply(conn)
				} else {
					respondIndex(conn)
				}
			case "POST":
				if fs[1] == "/apply" {
					respondPost(conn)
				}
			}
		}
		count++
	}
}
