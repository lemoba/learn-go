/*
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
error类型是一个接口类型，这是它的定义：

	type error interface {
		Error() string
	}
*/
package main

import "fmt"

// 定义divideError结构体
type DivideError struct {
	dividee int
	divider int
}

// 实现`error`接口
func (de DivideError) Error() string {
	strFormat := `
	Connot proceed, the divider is zero.
	dividee: %d,
	divider: 0,
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义`int`类型出发运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDivider,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}
func main() {
	// 正常情况
	if result, errMsg := Divide(100, 10); errMsg == "" {
		fmt.Println("100/10 = ", result)
	}

	// 当除数为零的时候会返回错误信息
	if _, errMsg := Divide(10, 0); errMsg != "" {
		fmt.Println("errorMsg is: ", errMsg)
	}
}
