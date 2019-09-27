package define

import (
	"net/http"
	"net/http/httputil"
)

type HH interface {
	GetBaidu(url string) string
	GetGoole()
}

type BuyTea struct {
	content string
}

func (t BuyTea)GetBaidu(url string) string{
	response, err := http.Get(url)
	defer response.Body.Close()
	if err!=nil{
		panic(err)
	}
	dumpResponse, e := httputil.DumpResponse(response, true)
	if e!=nil{
		panic(e)
	}
	return string(dumpResponse)
}