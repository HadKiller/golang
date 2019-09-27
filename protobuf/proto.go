package main

import (
	"ProtocSocket/example"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	_ "os"
)

func main() {
	//goproto:=&example.SearchRequest{
	//	Query:"天气",
	//	PageNumber:23,
	//	ResultPerPage:98}
	//marshal, e := proto.Marshal(goproto)
	//if e!=nil{
	//	fmt.Errorf("error")
	//}
	//ioutil.WriteFile("fab.txt",marshal,os.ModePerm)
	readFile()

}
func readFile()  {
	bytes, e := ioutil.ReadFile("fab.txt")
	serch:=&example.SearchRequest{}
	proto.Unmarshal(bytes,serch)
	fmt.Println(e)
	fmt.Println(serch.ResultPerPage,serch.PageNumber,serch.Query)

}
