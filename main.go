package main

import "fmt"

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
func main() {
	// coding
	var a, b = getSumAndSub(3, 2)
	fmt.Println(a, b)
}
