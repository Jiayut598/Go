// go切片
// 切片是拥有相同类型元素的可变长度的序列是基于数组类型做的一层封装 灵活支持自动扩容
// 切片拥有自己的长度和容量 len()求长度 cap()求切片的容量
package main

import "fmt"

func main(){
	/*
	var a []string
	var b = []int{}
	var c = []bool{false,true}
	// var d = []bool{false,true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)

	a := [5]int{1,2,3,4,5}
	s := a[1:3]
	fmt.Println(s)
	s2 := s[3:4]
	fmt.Println(cap(s2))
	// 完整的切片表达式
	// a[low : high : max]
	a := [5]int{1,2,3,4,5}
	t := a[1:3:5]
	fmt.Println(t)
	*/

	// 使用make()函数构造切片
	// make([]T, size, cap)
	// T切片的元素类型 size切片中元素的数量 cap切片的容量
	/*
	a := make([]int,2,10)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	*/
	// 切片的本质
	// 对底层数组的封装 包含了 底层数组的指针 切片的长度len() 切片的容量cap()
	// 判断数组是否为空 使用   len(s) == 0
	// 切片不能直接比较

	// 切片的赋值拷贝
	// 对一个切片的修改会影响另一个切片的内容，这点需要特别注意。
	/*
	s1 := make([]int,3)
	s2 := s1
	fmt.Println(s1)
	fmt.Println(s2)
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	*/

	 // 切片的遍历
	 /*
	 s := []int{1,2,3,4,5}
	 for i := 0;i <len(s); i++{
	 	fmt.Println(i,s[i])
	 }

	 for index,value := range s{
	 	fmt.Println(index,value)
	 }
	 */

	 // append方法为切片添加元素
	 /*
	 var s []int
	 s = append(s,1)
	 fmt.Println(s)
	 s = append(s,2,3,4)
	 fmt.Println(s)
	 s2 := []int{5,6}
	 fmt.Println(s2)
	 s = append(s, s2...) 
	 fmt.Println(s)

	 var numSlice []int
	 for i := 0; i < 10;i++{
	 	numSlice = append(numSlice,i)
	 	fmt.Println(numSlice)
	 }


	 // 一次追加多个元素
	 var citySlice []string
	 citySlice = append(citySlice,"北京")
	 fmt.Println(citySlice)
	 citySlice = append(citySlice,"上海","广州")
	 fmt.Println(citySlice)
	 // 追加切片
	 a := []string{"山西","河北"}
	 citySlice = append(citySlice,a...)
	 fmt.Println(citySlice)
	 */

	 // 使用copy函数复制切片
	 // 由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。
	 /*
	 a := []int{1,2,3,4,5}
	 b := a
	 fmt.Println(a)
	 fmt.Println(b)
	 b[0] = 100
	 fmt.Println(a)
	 fmt.Println(b)
	 // Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下：
	 // copy(destSlice, srcSlice []T)   srcSlice: 数据来源切片  destSlice: 目标切片
	 c := []int{1,2,3,4,5}
	 d := make([]int,5,5)
	 fmt.Println(c)
	 fmt.Println(d)
	 // d copy ---> c
	 copy(c,d)
	 fmt.Println(c)
	 fmt.Println(d)
	 */

	 // 切片中删除元素
	 /*
	 a := []int{1,2,3,4,5,6,7}
	 a = append(a[:2],a[3:]...)
	 fmt.Println(a)
	 fmt.Println(a[:2])
	 fmt.Println(a[3:])
	 */
	 var a = make([]string,5,10)
	 fmt.Println(a)
	 for i := 0; i < 10; i++{
	 	a = append(a,fmt.Sprintf("%v",i))
	 	fmt.Sprintf("%v",i)
	 }
	 fmt.Println(a)
}

