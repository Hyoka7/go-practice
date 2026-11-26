package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Connections map[net.Conn]string

type ServerInfo struct {
	connections Connections
	mutex       sync.Mutex
}

func (info *ServerInfo) broadcast(con net.Conn, msg string) {
	info.mutex.Lock()
	for c := range info.connections {
		if c != con {
			fmt.Fprintln(c, msg)
		}
	}
	info.mutex.Unlock()
}

func handle(conn net.Conn, info *ServerInfo) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var name string
	fmt.Fprint(conn, "enter your name:")
	ok := scanner.Scan()
	if !ok {
		return
	}
	name = scanner.Text()
	msg := fmt.Sprintf("%s is online", name)
	fmt.Println(msg)
	info.broadcast(conn, msg)
	fmt.Fprintln(conn, "Welcome to chat room,", name)
	info.mutex.Lock()
	info.connections[conn] = name
	info.mutex.Unlock()
	for scanner.Scan() {
		msg := fmt.Sprintf("%s:%s", name, scanner.Text())
		info.broadcast(conn, msg)
	}
	msg = fmt.Sprintf("%s is offline", name)
	fmt.Println(msg)
	info.broadcast(conn, msg)
	info.mutex.Lock()
	delete(info.connections, conn)
	info.mutex.Unlock()
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
	var info ServerInfo
	info.connections = make(Connections)
Server:
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			continue Server
		}
		go handle(conn, &info)
	}
}
