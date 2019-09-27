package main

import (
	"guangdu/my"
)

func main() {
	initMap := my.InitMap("src/guangdu/my/map.in")
	my.Walk(&initMap)

}

