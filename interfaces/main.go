package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	writeFile("fabb.txt")
	//fmt.Print( define.BuyTea{}.GetBaidu("http://www.baidu.com"))
}

func  writeFile(fileName string)  {
	file, e := os.Create(fileName)
	if e!=nil{
		panic(e)
	}
	defer file.Close()
	fd:=fabnacci()
	write:=bufio.NewWriter(file)
	for i:=0;i<30;i++  {
		fmt.Fprintln(write,fd())
	}
 defer 	write.Flush()
}
func fabnacci() func()  int{
	a,b:=1,1
	return func() int {
		a,b=b,a+b
		return a
	}
}
