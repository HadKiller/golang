package main

import (
	"fmt"
	"sync"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (t *Teacher) showT (){
	fmt.Println("show")
	t.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {

	m:=mutex{}
	m.m=make(map[string]string)
	m.write("hi","lou")
	m.read("hi")
	//runtime.GOMAXPROCS(1)
	//int_chan := make(chan int, 1)
	//string_chan := make(chan string, 1)
	//int_chan <- 1
	//string_chan <- "hello"
	//select {
	//case value := <-int_chan:
	//	fmt.Println(value)
	//case value := <-string_chan:
	//	panic(value)
	//}
}
func (mx *mutex) write(str1,str2 string)  {
	mx.Lock()
	defer  mx.Unlock()
	mx.m[str1]=str2

}
func (mx *mutex) read(str string)  {
	mx.Lock()

	defer mx.Unlock()
	if  i ,ok:= mx.m[str];ok{
		fmt.Print(i)
	}

}

type mutex struct {
	sync.Mutex
	m map[string]string
}
