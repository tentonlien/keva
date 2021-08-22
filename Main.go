package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

const port = 16379
const banner = "  _  __               \n | |/ /_____   ____ _ \n | ' // _ \\ \\ / / _` |\n | . \\  __/\\ V / (_| |\n |_|\\_\\___| \\_/ \\__,_|  v0.1.0\n"

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:" + strconv.Itoa(16379))
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()

	fmt.Println(banner)
	fmt.Println("Copyright 2021 Tenton Lien. All rights reserved.")
	fmt.Printf("Keva v0.1.0 listening on port " + strconv.Itoa(16379) + "...")
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		fmt.Println("A client connected:", tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}

}


func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("Disconnected:", ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		process(message)
		msg := time.Now().String() + conn.RemoteAddr().String() + "Server say hello!\n"
		b := []byte(msg)
		conn.Write(b)
	}
}
