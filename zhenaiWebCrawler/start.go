package main

import (
	"fmt"
	"time"
	"zhenaiWebCrawler/Save"
	"zhenaiWebCrawler/engine"
	"zhenaiWebCrawler/model"
	"zhenaiWebCrawler/parser"
)

func main() {

	start:=time.Now()
	e:=engine.SimpleEngine{
		UserSaveCha: Save.UserSaver(),
	}
	e.Run(model.Request{
		URL:       "http://www.zhenai.com/zhenghun",
		ParserFuc: parser.CityParser,
	})
	end:=time.Now()
	//start:=time.Now()
	//e:=engine.CurrentEngine{
	//	Schedule:   & engine.SimpleSchedule{},
	//	WorkGoCount: 10,
	//}
	//e.Run(model.Request{
	//		URL:       "http://www.zhenai.com/zhenghun",
	//		ParserFuc: parser.CityParser,
	//})
fmt.Println(end.Sub(start).Seconds())
}
