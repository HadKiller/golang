package engine

import (
	"fmt"
	fetch "zhenaiWebCrawler/fetcher"
	"zhenaiWebCrawler/model"
)

type SimpleEngine struct {
	UserSaveCha chan interface{}
}


func (e *SimpleEngine)Run(request ... model.Request){

	var requestQueue [] model.Request
	requestQueue=append(requestQueue, request ...)
	for len(requestQueue)>0{
		r:=requestQueue[0]
		requestQueue=requestQueue[1:]
		//合并fetcher和结果
		parserResult:=DoGowork(r)

		requestQueue=append(requestQueue, parserResult .NextRequest...)
		for _, v := range parserResult.Data {
		go func() {
			e.UserSaveCha<-v
			fmt.Printf("d")
		}()
			fmt.Printf("%v\n",v)
		}

	}
}
func DoGowork(r model.Request) *model.ParserResult{
	fetcher := fetch.Fetcher(r.URL)
	parserResult := r.ParserFuc(fetcher)
	return &parserResult
}
