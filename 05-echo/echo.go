package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()
	connIp := conn.RemoteAddr()
	fmt.Println(connIp, "is online")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Fprintln(conn, scanner.Text())
	}
	fmt.Println(connIp, "is offline")
	err := scanner.Err()
	if err != nil {
		fmt.Println("scanner error", err)
	}
}

func main() {
	ln, err := net.Listen("tcp", "localhost:4242")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	defer ln.Close()
Server:
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			continue Server
		}
		go handle(conn)
	}
}
