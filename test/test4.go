package test

import "fmt"

func Test1() {
	var n1 int32 = 12
	var n2 int64
	var n3 int8
	n2 = int64(n1) + 20
	n3 = int8(n1) + 20
	fmt.Println(n3, n2)
}
func Test2() {
	var n1 int32 = 12
	var n3 int8
	var n4 int8
	n4 = int8(n1) + 127 //【编译通过，但是 结果 不是127+12 ,按溢出处理】
	n3 = int8(n1) + 128 //【编译不过】
	fmt.Println(n4)
}
func Demo4() {
	// 数据类型转换
	var i int32 = 100
	var n1 float32 = float32(i)
	fmt.Println(n1)
	fmt.Println("练习---------")
	Test1()
	Test2()
}
