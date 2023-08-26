package test

import "fmt"

func Demo2() {
	var c1 byte = 'a'
	var c2 byte = '0' // 字符0
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	// 格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	// var c3 byte = '北'
	// fmt.Printf("c3=%c c3对应的码值=%d", c3, c3) // overflows byte
	var c4 int = '北'
	fmt.Printf("c4=%c c4对应的码值=%d\n", c4, c4)

	var n1 = 10 + 'a' // 10+97
	fmt.Println("n1=", n1)
}
