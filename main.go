package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/a-cordier/goose/proto/login"
)

func main() {
	conn, err := net.Dial("tcp", "server.slsknet.org:2242")
	if err != nil {
		log.Fatal("Unable to connect to server")
	}

	w := bufio.NewWriter(conn)
	loginMessage := login.Write("username", "password")
	w.Write(loginMessage)
	w.Flush()
	if err != nil {
		log.Fatal("login error")
	}
	res := bufio.NewReader(conn)
	response := login.Read(res)
	if response.OK() {
		success := response.(login.Success)
		fmt.Println(success.Greet, success.IP, success.Sum)
	}
}
