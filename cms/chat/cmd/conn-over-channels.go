package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

var clients = []net.Conn{}

// tcp server to initiate and listen using tcp port 8000
func serve() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	log.Println("Listening on port 8000 ... ")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConn(conn)
	}
}

// tcp connection handler
func handleConn(conn net.Conn) {
	// appends tcp connections through Conn interface into a slice clients
	clients = append(clients, conn)
	// input is a buffer io which creates a new cmd scanner on each cmd input
	input := bufio.NewScanner(conn)
	for input.Scan() {
		for _, c := range clients {
			c.Write([]byte("\t" + input.Text() + "Reply from server\n"))
		}
	}
	conn.Close()
}

// tcp client conn uses Dial method for the tcp client to connect to server
func conn() {
	c, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Println("Failed to connect:", err)

	}
	// go routine to copy from Stdin source to  dest conn, then copy from conn to Stdout
	go io.Copy(c, os.Stdin)
	io.Copy(os.Stdout, c)
}

// func main() {
// 	switch os.Args[1] {
// 	case "serve":
// 		serve()
// 	case "conn":
// 		conn()
// 	default:
// 		log.Println("Broken")
// 	}
// }
