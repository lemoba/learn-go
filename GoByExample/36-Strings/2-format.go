package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := &point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	fmt.Printf("struct2: %+v\n", p) // %+v的格式化输出内容将包括结构体的字段名

	fmt.Printf("struct3: %#v\n", p) // %#v 根据 Go 语法输出值，即会产生该值的源码片段

	fmt.Printf("type: %T\n", p) // 需要打印值的类型，使用 %T

	fmt.Printf("bool: %t\n", true) // 格式化格式化布尔值

	fmt.Printf("int: %d\n", 123) // 格式化整型数有多种方式，使用 %d 进行标准的十进制格式化

	fmt.Printf("bin: %b\n", 14) // 输出二进制表示形式

	fmt.Printf("char: %c\n", 65) // 输出给定整数的对应字符

	fmt.Printf("hex: %x\n", 456) // %x 提供了十六进制编码

	fmt.Printf("float1: %f\n", 78.9161)   // 使用 %f 进行最基本的十进制格式化
	fmt.Printf("float1: %.2f\n", 78.9141) // %.[num]f 可四舍五入保留指定位数

	fmt.Printf("float2: %e\n", 123400000.0) // 科学记数法表示形式
	fmt.Printf("float3: %E\n", 123400000.0) // 科学记数法表示形式

	fmt.Printf("str1: %s\n", "\"string\"") // 使用 %s 进行基本的字符串输出。

	fmt.Printf("str2: %q\n", "\"string\"") // 像 Go 源代码中那样带有双引号的输出

	fmt.Printf("str3: %x\n", "hex") // %x 输出使用 base-16 编码的字符串， 每个字节使用2个字符表示

	fmt.Printf("pointer: %p\n", &p) // 输出一个指针的值

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345) // %[Num]空值宽度 默认右对齐并用空格填充

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.451) // - 左对齐

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b") // - 右对齐

	s := fmt.Sprintf("sprintf: a %s", "string") // 我们已经看过 Printf 了， 它通过 os.Stdout 输出格式化的字符串。 Sprintf 则格式化并返回一个字符串而没有任何输出
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error") // Fprintf 来格式化并输出到 io.Writers 而不是 os.Stdout
}
