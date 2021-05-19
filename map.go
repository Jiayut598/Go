package main

import "fmt"

func main(){
	// go语言中提供映射关系的容器为map  内部使用散列表hash实现
	// map 是一种无序的基于key-value 的数据结构 go语言中map是引用类型 必须初始化才能使用
	// map[keytype]ValueType    keytype 表示键的类型 ValueType表示键对应值的类型
	// map类型遍历默认初始值nil 需要使用make函数来分配内存
	// make(map[KeyType]ValueType, [cap])

	// map基本使用
	/*
	scoreMap := make(map[string]int,8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 80
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of a:%T\n", scoreMap)
	*/
	/*
	// map也支持在声明的时候填充元素，例如：
	newMap := map[string]string{
		"username":"admin",
		"pwd":"admin",
	}
	fmt.Println(newMap)
	*/
	/*
	// 判断map中键是否存在
	// value,ok := map[key]
	v,ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	}else{
		fmt.Println("no")
	}

	// map遍历
	for k,v := range scoreMap{
		fmt.Println(k,v)
	}
	for k := range scoreMap{
		fmt.Println(k)
	}
	*/
	/*
	// 使用delete()函数删除键值对
	// delete(map,key)
	fmt.Println(scoreMap)
	delete(scoreMap,"李四")
	fmt.Println(scoreMap)
	*/
	// 按照制定顺序遍历map
	/*
	rand.Seed(time.Now().UnixNano())

	var nMap = make(map[string]int,200)

	for i := 0; i < 100; i++{
		key := fmt.Sprintf("stu%02d",i)
		value := rand.Intn(100)
		nMap[key] = value
	}

	var keys = make([]string,0,200)
	for key := range(nMap){
		keys = append(keys,key)
	}

	sort.Stings(keys)
	for _,key := range keys{
		fmt.Println(key,nMap[key])
	}
	*/
	// 元素为map类型的切片
	/*
	var mapSlice = make([]map[string]string,3)
	for index,value := range mapSlice{
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")

	mapSlice[0] = make(map[string]string,10)
	mapSlice[0]["name"] = "li1"
	mapSlice[0]["passwd"] = "li2"
	mapSlice[0]["adress"] = "li3"
	for index,value := range mapSlice{
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	*/
	// 值为切片类型的map
	/*
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
	*/

	// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。

}	