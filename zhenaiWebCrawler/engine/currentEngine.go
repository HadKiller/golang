package engine

import (
	"fmt"
	"zhenaiWebCrawler/model"
)

type CurrentEngine struct {
	        Schedule  Scheduler
	        WorkGoCount int
}

type Scheduler interface {
	Submit(m model.Request)

	ConfigureWork(chan model.Request)
}
func (c *CurrentEngine) Run(Seeds ... model.Request){

	for _, v := range Seeds {
		c.Schedule.Submit(v)
	}
	in:=make(chan model.Request)
	out:=make(chan model.ParserResult)
	c.Schedule.ConfigureWork(in)
	for i:=0;i<c.WorkGoCount ;i++  {
		CreateGoWork(in,out)
	}
	for{
		result:=<-out
		for _, v := range result.Data {
			fmt.Printf("%v",v)
		}

		for _, v := range result.NextRequest {
			c.Schedule.Submit(v)
		}
	}

}

func CreateGoWork(in chan model.Request, out chan model.ParserResult) {
	go func() {
		result:=<-in
		parseRe:=DoGowork(result)
		out<-*parseRe
	}()
}