package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	CreateServer()
}
func CreateServer(){
	tcpServer, _ := net.ResolveTCPAddr("tcp4", ":9999")
	listener, _ := net.ListenTCP("tcp", tcpServer)
	for{
		conn, e := listener.Accept()

		if e!=nil{
			fmt.Println(e)
			continue
		}
		conn.Write([]byte("Hi"))
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Minute))
	//conn.Write([]byte("Hi,Im jack"))
	request:=make([]byte,1024)
	for {
		len, err := conn.Read(request)
		if err!=nil{
			break
		}
		if len==0{
			fmt.Println("close Client")
			break
		}else {
			fmt.Println(string(request[:len])+"----form Client Msg")
			conn.Write([]byte("Hi,Im jack"))
			time.Sleep(time.Second)

		}
		request=make([]byte ,1024)
	}
}
