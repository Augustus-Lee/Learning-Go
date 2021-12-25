package main

import (
	"fmt"
	"reflect"
)

func add(num int) {
	num += 1

}

func realAdd(num *int)	{
	*num += 1

}


func main() {

	fmt.Println("hello test")


	var a int //声明
	a = 100   //赋值

	b := 200 //声明+赋值

	fmt.Println("a:", a, ",b:", b)


	str1 := "Golang"
	str2 := "Go语言"
	fmt.Println(reflect.TypeOf(str2[2]).Kind()) //打印某个变量类型，字符串以byte数组形式保存
	fmt.Println(str1[2], string(str1[2]))


	//将string转换成rune数组,可以正确处理中文
	str3 := "Go语言"
	runeArr := []rune(str3)
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind())
	fmt.Println(runeArr[2], string(runeArr[2]))
	fmt.Println("len(runeArr):", len(runeArr))

	fmt.Println("===============================")
	arr := [6]int{1,2,3,4,5,6}

	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}

	fmt.Println(arr)


	fmt.Println("===============================")

	//字典  仅声明
	m1 := make(map[string]int)

	//声明加初始化
	m2 := map[string]string {
		"Sam": "Male",
		"Alice": "Female",
	}

	m1["Tom"] = 22
	fmt.Println(m1)
	fmt.Println(m2)

	fmt.Println("===============================")

	num := 100
	add(num)
	fmt.Println(num)

	realAdd(&num)
	fmt.Println(num) 

	fmt.Println("===============================")

	//遍历字典
	m3 := map[string]string {
		"Sam": "Male",
		"Alice": "Female",
		"AA": "Male",
	}
	for key, val := range m3 {
		fmt.Println(key, val)
	}



}
