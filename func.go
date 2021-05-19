package main

import "fmt"

// 格式
/*
func 函数名(参数)(返回值){
	函数体
}
*/


/*
func intSum(x,y int)int{
	return x + y
}
func main(){
	sum1 := intSum(1,2)
	fmt.Println(sum1)
}
*/
// 可变参数
func intSum2(x...int)int{
	fmt.Println(x)
	sum := 0
	for _,v := range x{
		sum += v
	}
	return sum
}
func main(){
ret1 := intSum2()
ret2 := intSum2(10)
ret3 := intSum2(10, 20)
ret4 := intSum2(10, 20, 30)
fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60


}