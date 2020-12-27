package main

import "fmt"

func updateSlice(arr []int){
	arr[0]=100
}

func main(){

	arr := [...]int{0,1,2,3,4,5,6,7,8}

	s1 := arr[2:6]
	s2 := s1[3:5]

	//[2,3,4,5]
	fmt.Println(s1)
	//[5,6]
	//slice可以向后扩展，不可以向前扩展
	//s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)
	fmt.Println(s2)

	fmt.Println("========================")
	//slice本身没有数据，是对底层的array的一个view
	slice1 := arr[2:]
	slice2 := arr[:6]
	fmt.Println("arr = " ,arr)
	fmt.Println("slice1 = " ,slice1)
	fmt.Println("slice2 = " ,slice2)
	updateSlice(slice1)
	fmt.Println("after update arr = " ,arr)
	fmt.Println("after update slice1 = " ,slice1)
}