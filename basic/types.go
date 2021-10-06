package main

import (
	"fmt"
	"strconv"
)

//定义一个接口
type myInterface interface {
	sayHello() string
}

//定义一个数据类型
type myStruct struct {
	age int
}

//使用数据类型指针绑定方法
func (a *myStruct) sayHello() string {
	return "my age is "+ strconv.Itoa(a.age)
}

//定义一个函数，参数为一个自定义的数据类型
func printValueByStruct(arg myStruct){
	fmt.Printf("the age is %d \n",arg.age)
}

//定义一个函数，参数为一个接口
func printTheValue(arg myInterface) string{
	fmt.Printf("the value is %s \n",arg.sayHello())
	return arg.sayHello()
}

//定义一个函数，参数为动态个数的接口类型参数
func printAnyValue(args ...interface{}){
	//使用for range方法获取每一个接口
	for _,arg := range args{
		//使用.(type)方法查询接口的数据类型
		switch arg.(type) {
		case int:
			fmt.Println("the type is int")
		case string:
			fmt.Println("the type is string")
		case myStruct:/*是自定义数据类型*/
			//使用.(数据类型)进行强转
			var b myStruct = arg.(myStruct)
			fmt.Println("the type is myStruct, the function value is ", b.sayHello(),"the struct value is ", b.age/*调用数据结构的数据*/)
		case myInterface:/*是定义的接口数据类型*/
			fmt.Println("the type is myInterface interface, the function value is ", arg.(myInterface).sayHello()/*将接口强转到指定接口，并调用方法*/)
		}
	}
}

func main()  {
	//创建一个对象指针
	var structArg1 *myStruct= new(myStruct)
	//赋值
	structArg1.age = 27

	var structArg2 myStruct = myStruct{28} //创建一个对象，给对象赋值

	var interfaceArg1 interface {}=myStruct{29} //创建一个对象，将对象赋值后传给一个万能类型接口
	var interfaceArg2 interface{}=&myStruct{30}//创建一个对象，将对象指针传给一个万能类型接口


	printTheValue(structArg1) //V1会转换为myInterface接口被调用其中的方法
	printTheValue(interfaceArg2.(myInterface)) //万能接口interfaceArg2中放置的对象指针被强制转为myInterface接口调用
	printValueByStruct(*structArg1) //强制将structArg1的对象使用*显示传入函数，因为参数是对象
	printValueByStruct(interfaceArg1.(myStruct))//强制将万能接口a中放置的对象转换为对象传入函数，因为参数是对象

	printTheValue(&structArg2) //将对象的指针传入函数，golang将其转换为myInterface接口
	printAnyValue(structArg1/*传入一个指针，会同myInterface接口数据类型匹配*/,
		structArg2/*传入一个对象，会同myStruct数据类型匹配*/,
		*structArg1/*将指针显示为对象传入，会同myStruct数据类型匹配*/,
		&structArg2/*将对象的指针传入，会同myInterface接口数据类型匹配*/)

}