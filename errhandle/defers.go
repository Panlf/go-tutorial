package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
	defer如何使用
	Open/Close
	Lock/UnLock
	PrintHeader/PrintFooter


	panic
	停止当前函数执行
	一直向上返回，执行每一层defer
	如果没有遇见recover，程序退出

	recover
	仅在defer调用中使用
	获取panic的值
	如果无法处理，可重新panic
*/
func tryDefer(){
/*	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
*/
	for i:=0;i<100;i++{
		defer fmt.Println(i)
		if i==10{
			panic("print too many")
		}
	}
}


func writeFile(filename string){
	file,err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

}

func main(){
	tryDefer()
}
