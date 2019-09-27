package main

import (
	"fmt"
	"time"
)

func main() {

	for i:=0;i<300;i++{
		go func(i int) {
			fmt.Println("goroutine loop:", i)
		}(i)
	}
	time.Sleep(time.Microsecond)
}



