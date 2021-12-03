package main

import (
	"fmt"
	"strings"

)


func main() {
	fmt.Printf("Hello, World")
	var b = 1
	fmt.Println(b,"\n",b)
	str := "这里是 www\n.runoob\n.com"
	fmt.Println("-------- 原字符串 ----------")
	fmt.Println(str)
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println("-------- 去除空格与换行后 ----------")
	fmt.Println(max(1,2),"\n")
}
func max(num1,num2 int) int{
	var res int;
	if(num1 > num2){
		res  = num1
	}else {
		res = num2
	}
	return res
}
