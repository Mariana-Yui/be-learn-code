package main

import "fmt"

// 函数外定义还是需要, 可以使用()集中定义
var (
	aa     = 1
	bb, cc = 2, 3
	dd     = "Yui"
)

func variableZero() {
	var num int
	var str string
	// %s显示字符串 %q会在字符串两边加引号方便识别空串
	fmt.Printf("%d %q\n", num, str)
}

func variableInitialValue() {
	var num int = 3
	var str string = "Mariana"
	fmt.Println(num, str)
}

// 类型推导
func variableTypeDeduction() {
	var a, b = 3, 4
	var str = "Mariana"
	fmt.Println(a, b, str)
}

// 推荐的函数定义变量方式
func variableShorter() {
	// 只能在函数内这样定义
	a, b, c, d := 1, 2, 3, "Mariana"
	b = 4
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println("hello world")
	variableZero()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
