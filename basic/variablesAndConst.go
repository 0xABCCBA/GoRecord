package main

import (
	"fmt"
	"math"
)
//支持字符，字符串，布尔和数值 常量
const aConst string = "constant"

func main() {
	fmt.Println(aConst)
	//const可以出现在任何var语法可以出现的地方
	const n = 500000000
	//常数表达式可以执行任意精度的运算
	const v = 3e20 / n
	fmt.Println(v)
	//数值型常量是没有确定的类型的，直到它们被指定了一个类型
	fmt.Println(int64(v))
	//当上下文需要时，一个数可以被给定一个类型（包含变量赋值和函数调用）
	//比如下面的math.Sin，函数需要一个float64类型的参数
	fmt.Println(math.Sin(n))

	var a string = "initial"
	var b, c int = 1, 2
	var d = true
	var e int
	var f string
	g := "short"
	fmt.Println(a, b, c, d, e, f, g)
}
