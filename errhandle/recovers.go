package main

import (
	"fmt"
)

func tryRecover(){
	defer func(){
		r := recover()
		if err,ok := r.(error); ok{
			fmt.Println("Error occurred:",err)
		}else{
			panic(fmt.Sprintf("I don't know what to do : %v",r))
		}
	}()

	//a := 5 / 0
	//fmt.Println(a)
	//panic(errors.New("this is an error"))
	panic(12)
}

func main(){
	tryRecover()
}