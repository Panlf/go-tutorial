package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var(
	aa = 2
	bb = "kkk"
	cc = true
)

func variableZeroValue(){
	var a int //0
	var s string //""
	fmt.Printf("%d %q\n",a,s)
}

func variableInitialValue(){
	var a,b int = 3,4
	var s string= "123"
	fmt.Println(a,b,s)
}

func variableTypeDeduction(){
	var a,b,c,s = 3,4,true,"123"
	fmt.Println(a,b,c,s)
}

func variableShorter(){
	a,b,c,s := 3,4,true,"123"
	b = 5
	fmt.Println(a,b,c,s)
}

func euler(){
	fmt.Println(cmplx.Exp(1i*math.Pi)+1)
	fmt.Printf("%.3f\n",cmplx.Exp(1i*math.Pi)+1)
}

func triangle(){
	var a,b int = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a*+b*b)))
	fmt.Println(c)
}

func constant() {
	/*const fileName = "abc.txt"
	const a,b = 3,4*/
	const (
		fileName = "abc.txt"
		a,b = 3,4
	)
	var c int;
	c = int(math.Sqrt(a*a+b*b))
	fmt.Println(fileName,c)
}

func enums(){
	const(
		//iota 自增值
		cpp = iota
		//跳过
		_
		java
		python
		golang
	)

	const(
		b = 1 << (10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}


func main() {
	//常量定义枚举值
	enums()
	//常量定义
	constant()
	//数值计算
	euler()
	triangle()

	fmt.Println(aa,bb,cc)
 	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
