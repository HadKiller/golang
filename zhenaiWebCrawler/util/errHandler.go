package util

import "fmt"

func CheckErr(e error,msg string)  {
	if e!=nil{
		fmt.Println(msg)
	}
}
