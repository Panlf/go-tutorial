package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a,b int,op string) (int,error){
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		q,_ := div(a,b)
		return q,nil
	default:
		//panic("unsupported operation : "+op)
		return 0,fmt.Errorf("unsupported operation : %s",op)
	}
}

func div(a,b int) (int,int){
	return a/b,a%b
}

// 传入函数
func apply(op func(int,int) int, a,b int) int{
	p :=reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling funciton %s with args " +"(%d,%d)\n",opName,a,b)

	return op(a,b)
}

//可变参数
func sum(numbers ...int) int {
	s := 0
	for i:=range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a,b *int){
	*a,*b = *b,*a
}


func main(){
	fmt.Println(eval(3,2,"x"))


	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a),float64(b)))
	},2,3))

	fmt.Println(sum(2,3,1,3))

	a := 3
	b := 4

	fmt.Printf("a=%d,b=%d\n",a,b)

	swap(&a,&b)

	fmt.Printf("a=%d,b=%d",a,b)

}
