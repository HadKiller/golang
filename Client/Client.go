package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)
var gLocker sync.Mutex //全局锁
var gCondition *sync.Cond //全局i条件变量
func main() {
	tcpAddr, e := net.ResolveTCPAddr("tcp", ":9999")
	if checkErr(e,"Add"){
		return
	}

	conn, e := net.DialTCP("tcp", nil, tcpAddr)
	if checkErr(e,"dial"){
		return
	}
	fmt.Println("Connect to Server")
	gLocker.Lock()
	gCondition=sync.NewCond(&gLocker)
	conn.Write([]byte(" I am a Client！"))
	go clientConnHandler(conn)
	for{
		gCondition.Wait()
		break
	}
	gLocker.Unlock()
	fmt.Println("Client finish.")

}

func clientConnHandler(conn *net.TCPConn) {

	gLocker.Lock()
	defer gLocker.Unlock()
	defer conn.Close()
	request:=make([]byte,1024)
	for{
		len, e := conn.Read(request)
		if checkErr(e,"Read"){
			gCondition.Signal()
			break
		}
		if len==0{
			fmt.Println("Server connection close!");

			//条件变量同步通知
			gCondition.Signal();
			break;
		} else {
			//输出接收到的信息
			fmt.Println(string(request[:len]))

			//发送
			conn.Write([]byte("Hello !"));
		}

		request = make([]byte, 128);
	}
}

//错误处理函数
func checkErr(err error, extra string) bool {
	if err != nil {
		formatStr := " Err : %s\n";
		if extra != "" {
			formatStr = extra + formatStr;
		}

		fmt.Fprintf(os.Stderr, formatStr, err.Error());
		return true;
	}

	return false;
}
