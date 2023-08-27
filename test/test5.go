package test

import (
	"fmt"
	"strconv"
)

// 基本数据类型和string转换
func Test4() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string // 空的str

	//使用第一种方式来转换 fmt.Sprintf方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)
	fmt.Println("------------------")
	//第二种方式 strconv 函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	// strconv.FormatFloat(num4, 'f', 10, 64)
	// 说明： 'f' 格式 10：表示小数位保留10位 64 :表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	//strconv包中有一个函数Itoa
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str=%q\n", str, str)

}
func Test5() {
	var i int = 10
	fmt.Println("i的地址：", &i)
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)
}

func Test6() {
	// 1)写一个程序,获取一个int变量num的地址,并显示到终端
	// 2)将num的地址赋给指针ptr,并通过ptr去修改num的值.
	var num int
	fmt.Printf("num的地址:%v\n", &num)
	var ptr *int = &num
	*ptr = 11
	fmt.Println("num的值: ", num)

}
func Demo5() {
	Test6()
}
