package fetch

import (
	"io/ioutil"
	"net/http"
	"zhenaiWebCrawler/util"
)

func  DoWork()  {

}
//抓取页面信息
func Fetcher(url string)*[]byte{
	resp, err := http.Get(url)
	util.CheckErr(err,"httpGet fail:"+string(resp.StatusCode)+
		"check net ,address！")
	defer resp.Body.Close()
	resultBytes, err := ioutil.ReadAll(resp.Body)
	util.CheckErr(err,"read page fail！")
	return &resultBytes
}
