package Save

import "fmt"


func UserSaver() chan interface{}{
	out:=make(chan interface{})
	go func() {
		for {
			user:=<-out
			fmt.Printf("savggggggge%v",user)
		}

	}()
	return out
}