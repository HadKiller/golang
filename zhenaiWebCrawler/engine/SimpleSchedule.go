package engine

import "zhenaiWebCrawler/model"

type SimpleSchedule struct {

	WorkChan chan  model.Request
}

func (s *SimpleSchedule) ConfigureWork(c chan model.Request) {
	s.WorkChan=c
}

func (s *SimpleSchedule) Submit(r model.Request) {
	s.WorkChan<-r
}
