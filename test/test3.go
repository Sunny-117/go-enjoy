package test

import "fmt"

func Demo3() {
	var address string = "12312"
	fmt.Println(address)
	// address[5] = 'X' //不可修改
	str3 := `
	package: main
	import (
		"fmt"
		"unsafe"
	)
	func main() {
		fmt.Println(12)
	}
	`
	fmt.Println(str3)
	fmt.Println("abc\nabc")

	var str = "hello" + " world"
	str += " haha"
	fmt.Println(str)
	var z int
	fmt.Println(z) //默认值
}
