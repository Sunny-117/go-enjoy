package test

import "fmt"

func Demo1() {
	var i int
	fmt.Println("i=", i)
	var num float64 = 10.11111111111
	fmt.Println("num=", num)
	name := "tom"
	fmt.Println("name=", name)
	n1, name, n3 := 100, "tom~", 888
	fmt.Println(n1, name, n3)
	var (
		x     = 300
		n4    = 900
		name2 = "marry"
	)
	fmt.Println(x, n4, name2)
	var j int = 10
	j = 20
	fmt.Println(j)
	var bz = 100
	fmt.Printf("n1的类型%T \n", bz)
}
