package main

import (
	"fmt"
	"net"
)

func startUDPServer (port int) {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: port})
	check(err)
	fmt.Println ("UDP Server start at port: ", port)
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := udpConn.ReadFromUDP(data)
		check(err)
		fmt.Printf("[%s] %s\n", remoteAddr, data[:n])
		_, err = udpConn.WriteToUDP(data[:n], remoteAddr)
		check(err)
	}

}