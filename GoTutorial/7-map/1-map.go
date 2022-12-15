// Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
// Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

// 1.定义Map
// 可以使用内建函数make 也可以使用map 关键字来定义Map:

/* 1.1 声明变量, 默认map是nil */
// var map_variable map[key_data_type]value_data_type

/* 1.2 使用make函数 */
// map_variable := make(map[key_data_type]value_data_type)

// 2. 使用

// 声明
// var countryCapitalMap map[string]string

package main

import "fmt"

func main() {
	countryCapitalMap := make(map[string]string)

	// 赋值
	countryCapitalMap["China"] = "Beijing"
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	// 遍历
	PrintMap(countryCapitalMap)

	// 查看元素在集合集合中是否存储在
	capital, ok := countryCapitalMap["China"]

	if ok {
		fmt.Printf("China的首都是%s\n", capital)
	} else {
		fmt.Println("China的首都不存在")
	}

	delete(countryCapitalMap, "India")

	fmt.Println("删除后的元素")

	PrintMap(countryCapitalMap)

}

func PrintMap(countryMCapitalMap map[string]string) {
	for country, city := range countryMCapitalMap {
		fmt.Println(country, "首都是", city)
	}
}
