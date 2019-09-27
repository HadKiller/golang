package main

import (
	"fmt"
	"time"
)

type workChannel struct {
	send chan int
	done chan  bool
}

func dowork(id int ,char workChannel)  {
	fmt.Printf("work id: %d  -result:%c",id,<-char.send)
	go func() {
		char.done<-true
	}()
}
func createChan(chanId int) workChannel{
	work:=workChannel{
		send: make(chan int),
		done: make(chan bool),
	}
	go dowork(chanId,work)
	return work
}
func cha()  {
	ch1:=make(chan  int)

	go func() {
		var n int
		for{
			n=<-ch1
			fmt.Print(n)
		}
	}()
	ch1<-12
	ch1<-90
	time.Sleep(time.Second)
}
func main()  {
	var chaArray [10]workChannel

	for i:=0;i<10;i++{
		chaArray[i]=createChan(i)
	}

	for i:=0;i<10;i++{
		chaArray[i].send<-i+'d'
	}
	for i:=0;i<10;i++{
		chaArray[i].send<-i+'A'
	}
	for _,v:=range chaArray{
		<-v.done
		<-v.done
	}
}