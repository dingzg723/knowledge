package main

import "fmt"

func main() {
	var a int
	/* 声明指针变量 */
	var b *int

	/* 声明实际变量 */
	a = 5
	fmt.Println("a=", a)

	// 将a变量的实际内存地址赋值给b
	b = &a
	fmt.Println("b=", b)

	// *指针，表示使用指针访问值
	fmt.Println("b值为", *b)

	// 修改指针指向的值，则a变量被修改
	*b = 6
	fmt.Println("a=", a)

}
