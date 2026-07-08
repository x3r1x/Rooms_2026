package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//scp -i "C:\Users\maxim\.ssh\id_ed25519.pub" server yc-user@84.201.159.214:/home/yc-user/
	connection, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		if err := connection.Close(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	buff := make([]byte, 1024)
	for {
		n, remoteAdr, err := connection.ReadFromUDP(buff)
		if err != nil {
			fmt.Println(err)
			continue
		}
		msg := string(buff[0:n])

		fmt.Println("Получено от", remoteAdr, ":", msg)
		_, err = connection.WriteToUDP([]byte(msg), remoteAdr)
		if err != nil {
			fmt.Println(err)
		}
	}
}
