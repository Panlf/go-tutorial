package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	//写入信道
	done <- true
}

func numAdd(number int, add chan int){
	if number > 0 {
		number += 10
	}
	add <- number
}

func numCalc(number int,calc chan int){
	if number != 0 {
		number *= 2
	}
	calc <- number
}


func main() {
	done := make(chan bool)
	go hello(done)
	//读取管道
	fmt.Println(<-done)
	fmt.Println("main function")


	add := make(chan int)
	calc := make(chan int)

	number := 15

	go numAdd(number,add)
	go numCalc(number,calc)

	addResult,calcResult := <-add,<-calc

	fmt.Println("final result ", addResult+calcResult)


}