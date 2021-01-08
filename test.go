// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"net"
// 	"strconv"
// 	"time"
// )

// const serverAddress = "localhost:1337"

// func main() {
// 	fmt.Println("UDP client")

// 	ServerAddr, err := net.ResolveUDPAddr("udp", serverAddress)
// 	if err != nil {
// 		panic(err)
// 		fmt.Println("panic")
// 	}
// 	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
// 	if err != nil {
// 		panic(err)
// 	}

// 	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer Conn.Close()

// 	s1 := rand.NewSource(time.Now().UnixNano())
// 	r1 := rand.New(s1)

// 	ticker := time.NewTicker(3 * time.Second)

// 	go func() {
// 		for range ticker.C {
// 			msg := strconv.Itoa(r1.Intn(1000))
// 			buf := []byte(msg)

// 			fmt.Println("Sending > " + string(buf))

// 			_, err := Conn.Write(buf)

// 			if err != nil {
// 				fmt.Println(msg, err)
// 			}

// 			buf = make([]byte, 1024)
// 			n, _, err := Conn.ReadFromUDP(buf)
// 			if err != nil {
// 				fmt.Println("Error on receiving:", err)
// 			} else {
// 				fmt.Println("Recieved message from the server:", string(buf[0:n]))
// 			}
// 		}
// 	}()

// 	select {}

// }
