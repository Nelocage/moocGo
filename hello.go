package main

import (
	"fmt"
	"io/ioutil"
)

//并不是全局变量，而是包内的全局变量
var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variable() {
	var a int
	var s string
	var c, d int = 3, 4
	e := 5 //有var就没有引号，没有var就有引号
	fmt.Printf("%d,%q,%d,%d\n", a, s, c, d, e)
	//类型转换，go只有强制类型转换，没有隐式类型转换
	//常量可以规定类型，也可以不规定
	const filename = "abc.txt"
	const filename1 string = "abc.txt"

	//使用常量，定义枚举类型
	const (
		cpp        = 0
		java       = 1
		python     = 2
		golang     = 3
		javascript = iota
	)
	//iota可以参与运算  B KB MB GB
	const (
		b = 1 << (10 * iota) //iota可以理解为一个自增值的种子
		kb
		mb
		gb
		tb
		pb
	)

	/*定义变量
	变量类型写在变量名后
	编译器可推测变量类型
	没有char,只有rune*/

}

func ifelse() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else { //必须在同一行，我草
		fmt.Println(contents)
	}

	//go 的if和for语句都不带括号
}

//go的函数
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	default:
		panic("不支持的运算")
	}
}

//函数可以返回多个值
func div(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	fmt.Println("hello,world")
	variable()

	ifelse()

}

//rune就是go中的char
