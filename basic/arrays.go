package main

import "fmt"

func printArray(arr [3]int){
	for i,v := range  arr {
		fmt.Println(i,v)
	}
}

func editArray(arr *[3]int){
	arr[0] = 100
}

func main(){
	var arr1 [5]int
	arr2 := [3]int{1,7,2}
	arr3 := [...]int{2,4,6,7}

	//4行5列
	var grid [4][5]int

	fmt.Println(arr1,arr2,arr3,grid)

	printArray(arr2)

	fmt.Println("============================")

	arr4 := [3]int{2,3,6}
	printArray(arr4)
	editArray(&arr4)
	printArray(arr4)
}


