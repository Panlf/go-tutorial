package main

import "fmt"

//寻找最长不含有重复字符的子串
func lengthOfNonRepeatingSubStr(s string) int{
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i,ch := range []rune(s){
		lastI,ok := lastOccurred[ch]
		if ok && lastI >=  start{
			start = lastI + 1
		}
		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main(){
	length := lengthOfNonRepeatingSubStr("bbbCCAAS")
	fmt.Println(length)

	length = lengthOfNonRepeatingSubStr("中国心心")
	fmt.Println(length)
}
