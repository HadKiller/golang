package my

import (
	"fmt"
	"os"
)


type Ponit struct {
	I,j int
}
func InitMap(path string)  [][]int{
	file, e := os.Open(path)
	if e!=nil{
		panic(e)
	}


	var row,column int
	fmt.Fscanf(file,"%d %d",&row,&column)
	mapSlice:=make([][]int,row)
	for i :=range mapSlice{
		mapSlice[i]=make([]int,column)
		for j:=range mapSlice[i]{
			fmt.Fscanf(file,"%d",&mapSlice[i][j])
		}
	}
	return  mapSlice
}

func Walk(initMap *[][]int)  {

}