package main

import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {

	config := LoadConfiguration()
	fmt.Println(config.Port)
	fmt.Println(config.Technology)

	for i := 0; i < len(config.Redirect); i++ {
		fmt.Println(config.Redirect[i])
	}

	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: config.Port,
		IP:   net.ParseIP("0.0.0.0"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(ser, remoteaddr)

		for i := 0; i < len(config.Redirect); i++ {
			go sendString(string(p), config.Redirect[i])
		}
	}
}
