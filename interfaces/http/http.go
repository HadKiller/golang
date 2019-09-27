package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func getClient(){
	//response, err := http.Get("http://www.baidu.com")
	newRequest, newerr := http.NewRequest(http.MethodGet, "http://www.bilibili.com", nil)
	if newerr!=nil{
		panic(newerr)
	}
	newRequest.Header.Add("User-Agent",
		" Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
	response, e := http.DefaultClient.Do(newRequest)
	if e!=nil{
		panic(e)
	}
	dumpResponse, newerr := httputil.DumpResponse(response, true)
	fmt.Printf("%s",dumpResponse)
	//if err!=nil{
	//	panic(err)
	//}
	//defer  response.Body.Close()
	//dumpRequest, e := httputil.DumpResponse(response, true)
	//if e!=nil{
	//	panic(e)
	//}
	//fmt.Printf("%s",dumpRequest)

}
func main() {
	getClient()
}
