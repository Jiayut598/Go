package main

import "fmt"

func main(){
	/*
	var testArray [3]int
	var numArray = [...]int{1,2,3,4,5}
	var cityArray = [...]string{"北京","天津","南京"}
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)

	for i := 0;i < len(numArray); i++{
		fmt.Println(numArray[i])
	}
	for index,value := range cityArray{
		fmt.Println(index,value)
	}
	a := [4][2]int{
		{1,2},
		{3,4},
		{5,6},
		{7,8},
	}
	fmt.Println(a[3][1])
	for _,v1 := range(a){
		for _,v2 := range(v1){
			fmt.Println(v2)
		}
		fmt.Println()
	}
	*/
	// 求数组1，3，5，7，8的和
	/*
	var a = [5]int{1,3,5,7,8}
	sum := 0
	for i := 0;i < len(a);i++{
		sum += a[i]
	}
	fmt.Println(sum)
	*/
	// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	var a = [5]int{1,3,5,7,8}
	sum := 8
	for i := 0;i < len(a);i++{
		for j := i;j < len(a);j++{
			if a[i]+a[j] == sum{
				fmt.Printf("%d,%d\n",i,j)
			}
		}
	}
}
