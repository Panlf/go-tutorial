package main

import "fmt"


func changeValue(val *int) {
	*val = 100
}

func modifyArr(arr *[3]int) {
	//a[x] 是 (*a)[x] 的简写形式，因此上面代码中的 (*arr)[0] 可以替换为 arr[0]
	//(*arr)[0] = 90
	arr[0] = 90
}

func modifySlices(sls []int) {
	sls[0] = 90
}

func main(){

	a := 12

	var b *int = &a

	fmt.Printf("Type of b is %T\n", b)

	fmt.Println("b = ",b)

	var c *int

	fmt.Println("c = ",c)


	var m = 10

	fmt.Println("before m = ",m)

	changeValue(&m)

	fmt.Println("after m = ",m)


	arr := [3]int{89, 90, 91}
	modifyArr(&arr)
	fmt.Println(arr)


	slices := [3]int{89, 90, 91}
	modifySlices(slices[:])
	fmt.Println(slices)
}