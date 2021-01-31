package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"

	"github.com/davebehr1/goUDP/gameObjects"
)

type server struct {
	running    bool
	serverPort string
	// activeArenas map[string]*gameObjects.Arena
	activeArenas []*gameObjects.Arena
}

func (s *server) start() {
	s.running = true
}

func (s *server) run() {

	if s.running {
		ServerAddr, err := net.ResolveUDPAddr("udp", s.serverPort)

		if err != nil {
			panic(err)
		}

		ServerConn, err := net.ListenUDP("udp", ServerAddr)

		if err != nil {
			panic(err)
		}

		defer ServerConn.Close()

		var message gameObjects.Packet

		go func() {
			for {
				buf := make([]byte, 4096)
				n, addr, err := ServerConn.ReadFromUDP(buf)
				buffer := bytes.NewBuffer(buf[:n])
				decoder := gob.NewDecoder(buffer)
				err = decoder.Decode(&message)
				fmt.Println("Recieved", message.Payload, "from", addr)

				if err != nil {
					fmt.Println("Error:", err)
				}

				buf = []byte("OK")
				ServerConn.WriteToUDP(buf, addr)
			}
		}()

		fmt.Println("Waiting for clients to connect.Server Port " + s.serverPort)

		s.addNewArena()

		select {}
	}
}

func (s *server) addNewArena() {
	arena := &gameObjects.Arena{}
	s.activeArenas = append(s.activeArenas, arena)
}

func newServer(ServerPort string) *server {
	server := &server{
		false,
		ServerPort,
		make([]*gameObjects.Arena, 1),
	}

	return server
}

func main() {

	server := newServer(":1337")

	server.start()
	server.run()
	fmt.Println("A basic udp server")

	// ServerAddr, err := net.ResolveUDPAddr("udp", serverPort)
	// if err != nil {
	// 	panic(err)
	// }

	// ServerConn, err := net.ListenUDP("udp", ServerAddr)
	// if err != nil {
	// 	panic(err)
	// }

	// defer ServerConn.Close()

	// buf := make([]byte, 2048)

	// fmt.Println("Waiting for clients to connect.Server Port " + serverPort)

	// select {}
}
