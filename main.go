package main

import (
	"fmt"
	"test/calc"
)

func main() {
	fmt.Println("running...")
	// 写法1
	// var result1 int = calc.Plus(1, 2)
	// var result2 int = calc.Multiply(4, 2)
	// 写法2
	result1 := calc.Plus(1, 2) // int自动推断，类似ts的类型推断
	result2 := calc.Multiply(4, 2)
	fmt.Println("result1: ", result1)
	fmt.Println("result2: ", result2)
	calc.PrintOver()
}
