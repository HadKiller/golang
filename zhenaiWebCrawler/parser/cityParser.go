package parser

import (
	"encoding/json"
	"regexp"
	"zhenaiWebCrawler/model"
	"zhenaiWebCrawler/util"
)
var(
	cityFilter = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]*)"[^>]*>([\s\S]*?)</a>`)
	userFilter=regexp.MustCompile(`{"age":[^}]*"workCity":"[\s\S]*?"}`)
)

	//<a\b[^>]+\bhref="([^"]*)"[^>]*>([\s\S]*?)</a>
	//<a target="_blank" href="http://www.zhenai.com/zhenghun/beijing" data-v-7e67c21c>北京</a>


func CityParser(pageResult *[]byte) model.ParserResult {
	parseRe:=model.ParserResult{}
	cityData:=pageResult
	allSubmatch := cityFilter.FindAllSubmatch(*cityData, -1)
	for _,v:= range allSubmatch{
		parseRe.Data=append(parseRe.Data, string(v[2]))
		parseRe.NextRequest=append(parseRe.NextRequest,model.Request{
			URL:       string(v[1]),
			ParserFuc: CityInParser,
		})
	}
	return  parseRe
}

func CityInParser(data *[]byte) model.ParserResult{

	cityInData:=model.ParserResult{}
	findAll := userFilter.FindAll(*data, -1)
	var nu model.User
	for _, v := range findAll {
		err := json.Unmarshal(v, &nu)
		util.CheckErr(err,"反序列化出错")
		cityInData.Data=append(cityInData.Data, nu)
		cityInData.NextRequest=append(cityInData.NextRequest,model.Request{
			URL: nu.AvatarURL,
			ParserFuc: Nil,
		})
	}
	return cityInData
}

func Nil(*[]byte)model.ParserResult  {
	//TODO
	return model.ParserResult{
		NextRequest: nil,
		Data:        nil,
	}
}