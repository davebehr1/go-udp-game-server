package main

import (
	"fmt"
	"net"
)

const serverPort = ":1337"

func main() {
	fmt.Println("A basic udp server")

	ServerAddr, err := net.ResolveUDPAddr("udp", serverPort)
	if err != nil {
		panic(err)
	}

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		panic(err)
	}

	defer ServerConn.Close()

	buf := make([]byte, 2048)

	go func() {
		for {
			n, addr, err := ServerConn.ReadFromUDP(buf)
			fmt.Println("Recieved", string(buf[0:n]), "from", addr)

			if err != nil {
				fmt.Println("Error:", err)
			}

			buf = []byte("OK")
			ServerConn.WriteToUDP(buf, addr)
		}
	}()

	fmt.Println("Waiting for clients to connect.Server Port " + serverPort)

	select {}
}
