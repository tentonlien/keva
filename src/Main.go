package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
)

const port = 16379
const banner = "   ____  __  __    _  __               \n" +
	"  / /\\ \\ \\ \\ \\ \\  | |/ /_____   ____ _ \n" +
	" / /  \\ \\ \\ \\ \\ \\ | ' // _ \\ \\ / / _` |\n" +
	" \\ \\  / / / / / / | . \\  __/\\ V / (_| |\n" +
	"  \\_\\/_/ /_/ /_/  |_|\\_\\___| \\_/ \\__,_|  v0.1.0\n"


func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:" + strconv.Itoa(16379))
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()

	fmt.Println(banner)
	fmt.Println("Copyright 2021 Tenton Lien. All rights reserved.")
	fmt.Println("Keva v0.1.0 listening on port " + strconv.Itoa(16379) + "...")
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

	var cmd []string
	count := 0
	for {
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			fmt.Println("Error:", err.Error())
			break
		}
		if len(message) > 2 {
			message = message[0: len(message) - 2]
		}
		fmt.Println("message:", message)
		cmd = append(cmd, message)
		if message[0] == '*' {
			length, err := strconv.Atoi(message[1:])
			fmt.Println("length:", length)
			if err != nil {
				fmt.Println("Error", err.Error())
			}
			count = length
		} else if message[0] == '$' {
			count --
		}
		if count == 0 && message[0] != '$' {
			fmt.Println("CMD:", cmd)
			process(cmd)
			b := []byte("+OK\r\n")
			conn.Write(b)
			cmd = nil
		}
	}

	//msg := time.Now().String() + conn.RemoteAddr().String() + "Server say hello!\n"
	//b := []byte(msg)
	//conn.Write(b)
}
