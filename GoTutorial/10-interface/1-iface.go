// Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

/* 定义接口
type interface_name interface {
	method_name1 [return_type]
	method_name2 [return_type]
	...
	method_namen [return_type]
}
*/
/* 定义结构体
type struct_name struct {
	// variables
}
*/

/*
实现接口方法

	func (struct_name_variable struct_name) method_name1 [return_type] {
		// 方法实现
	}
*/
package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokia NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type XiaoMiPhone struct {
}

func (xiaomi XiaoMiPhone) call() {
	fmt.Println("I am XiaoMi, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(XiaoMiPhone)
	phone.call()
}
