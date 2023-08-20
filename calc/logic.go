package calc

// 1.func->function(函数)的声明
// 2. Plus -> 函数名称(需要公共调用的方法需要大坨峰) GetName
// 3. 类型定义->Go语言中类型定义放在变量后面
// 4. a与b都是函数参数,在调用(Plus(1,2))时,传入实际参数
// 5. 返回值类型定义在参数括号后面
// 6. go不需要打分号
// 7. 一个包内部调用不需要包引导，可以直接调用
// 8. 不同包之间调用方法需要有包名引导

func Plus(a int, b int) int {
	PrintOver()
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}
