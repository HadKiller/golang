package main

import (
	"fmt"
	"sync"
)

func main() {

	wg:=sync.WaitGroup{}
	//cha:=make(chan int,23)
	//wg.Add(2)
	//go printNum(cha,&wg)
	//go printStr(cha,&wg)
	//wg.Wait()
	//fmt.Print("\n")
	//paseStudent := pase_student()
	//fmt.Print(paseStudent)
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()

}



type student struct {
	Name string
	Age  int
}

func pase_student() map[string] *student{
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	return  m
}

func printNum(cha chan int,group *sync.WaitGroup){
	defer group.Done()
	for i:=1;i<6;i++{

		fmt.Printf("%d%d",2*i-1,2*i)
		cha<-1
		<-cha
	}
}

func printStr(cha chan int,group *sync.WaitGroup)  {
	defer  group.Done()
	for i:=0;i<5;i++{
		<-cha
		fmt.Printf("%c%c",65+i*2,66+i*2)
		cha<-1
	}


}
