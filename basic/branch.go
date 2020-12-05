package main

import (
	"fmt"
	"io/ioutil"
)


func grade(score int) string{
	g := ""
	switch  {
	case score<60:
		g="F"
	case score<80:
		g="C"
	case score<90:
		g="B"
	case score<=100:
		g="A"
	default:
		panic(fmt.Sprintf("Wrong score：%d",score))
	}
	return g
}

func main(){
	result :=grade(100)
	fmt.Println(result)

	//使用绝对路径
	const fileName = "E:\\CodeResource\\GoProject\\go-tutorial\\basic\\test.txt"

	/*content,err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}*/

	if content,err := ioutil.ReadFile(fileName);err == nil {
		fmt.Printf("%s\n", content)
	}else{
		fmt.Println(err)
	}

}
