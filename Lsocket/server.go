package main

import (
	"Lsocket/hero"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"time"
)

var timeTemplate1 = "2006-01-02 15:04:05" //常规类型
type endPoint struct {
	//protocol tcp tcp 4/6
	port, protocol string
}
func main() {
	address:=endPoint{
		port:     ":7788",
		protocol: "tcp4",
	}

	//获得监听
	Server:=CreateServer(address)
	ReceiveClient(*Server)
}

func CreateServer(p endPoint) *net.TCPListener{
	resolveTCPAddr, e := net.ResolveTCPAddr(p.protocol, p.port)
	CheckErr(e,"创建tcpAddress")
	tcpListener, e := net.ListenTCP(p.protocol, resolveTCPAddr)
	CheckErr(e,"监听地址")
	//tcpListener.SetDeadline()
	return tcpListener
}

func ReceiveClient(server net.TCPListener)  {
	log.Print("创建服务器。。。")
	for{
		tcpConn, e := server.AcceptTCP()
		CheckErr(e,"客户端连接")
		log.Println("客户端连接成功---")
		tcpConn.Write([]byte("I Am Server！"))
		weapons:=[]string {"水印帽子","无尽之刃","复活甲","耐克"}

		data:=hero.HeroAttack{
			Name:                "旋风斩",
			Damage:              234234,
			Weapons:             weapons,
		}
		bytes, e := proto.Marshal(&data)
		CheckErr(e,"序列化")
		i, e := tcpConn.Write(bytes)
		CheckErr(e,"向客户端发送问题")
		log.Print(i,"------------G-----------")
		go DoWork(tcpConn )
	}
}

func DoWork(conn *net.TCPConn) {
	defer conn.Close()
	receive:=make([]byte,1024)
	for{
		len, e := conn.Read(receive)
		CheckErr(e,"接受消息")
		if len==0{
			log.Print("客户端关闭")
			break
		}else {
			//fmt.Print("未经序列化：",string(receive[:len]),"--------",time.Now())
			heros:=hero.HeroAttack{}
			e := proto.Unmarshal(receive[:len], &heros)
			CheckErr(e,"protobuf反序列化")
			fmt.Printf("%v\n",heros)
			fmt.Println("收到客户端消息：","--------",time.Now())
			time.Sleep(time.Second)
			conn.Write([]byte("服务器收到消息了------------"+time.Now().Format(timeTemplate1)))
		}
	}

}
func CheckErr(e error,msg string)  {
	if e!=nil{
		log.Printf("%v%s",e,"------------"+msg+"出现错误！")
		return
	}
}