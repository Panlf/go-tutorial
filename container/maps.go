package main

import "fmt"

func main(){
	m := map[string]string {
		"name":"mouse",
		"course":"golang",
		"quality":"Good",
	}

	//创建
	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m,m2,m3)

	for k,v := range m {
		fmt.Println(k,v)
	}

	courseName := m["course"]
	fmt.Println(courseName)


	if name,ok := m["names"];ok {
		fmt.Println(name)
	}else {
		fmt.Println("key is not exist")
	}
}
