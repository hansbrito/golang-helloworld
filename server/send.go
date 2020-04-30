package main

import (
	"bufio"
	"fmt"
	"net"
)

func sendString(text string, host string) {
	fmt.Printf("Tentando enviar para o host:%s\n", host)

	p := make([]byte, 2048)
	conn, err := net.Dial("udp", host)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, text)
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
