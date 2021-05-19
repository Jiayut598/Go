package main

import (
	"fmt"
)


func main(){
	/*
	// 数组
	var testArray = [...]int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < len(testArray); i++{
		fmt.Println(testArray[i])
	}
	a := [...]string{"1","2","3"}
	for i := 0; i < len(a); i++{
		fmt.Println(a[i])
	}
	// 求1，2，3，4，5，6，7，8，9，10的和
	var testArray = [...]int{1,2,3,4,5,6,7,8,9,10}
	sum := 0
	for i := 0; i < len(testArray); i++{
		fmt.Println(testArray[i])
		sum += testArray[i]
	}
	fmt.Println(sum)
	// 找出指定下标
	var testArray = [...]int{1,2,3,4,5,6,7,8,9,10}
	sum := 20
	for i := 0; i < len(testArray); i++{
		for j := 0; j < len(testArray); j++{
			if testArray[i] + testArray[j] == sum{
				fmt.Println(i,j)
				fmt.Printf("%d,%d\n",i,j)

			}
		}
	}
	*/

	// 切片
	// var name []T
	/*
	var a []string
    var b = [5]int{1,2,3,4,5}
    s := b[1:3]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s)	
	*/

	// map
	// make(map[keytype]valuetype,[cap])
	/*
	var sorceMap = make(map[string]string,5)
	sorceMap["张三"] = "男"
	sorceMap["李四"] = "男"
	fmt.Println(sorceMap)
	*/

	var mapSlice = make([]map[string]string,1)
	mapSlice[0] = make(map[string]string)
	mapSlice[0]["name"] = "zhangsan"
	mapSlice[0]["sex"] = "男"
	fmt.Println(mapSlice)

}