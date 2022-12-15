package main

import (
	"encoding/json"
	"fmt"
)

type FamilMember struct {
	Name    string
	Age     int
	Parents []string
}

func main() {
	var b []byte = []byte(`{"Name": "Wednesday", "Age": 1, "Parents": ["Gomez", "Morticia"]}`)

	var f interface{}

	err := json.Unmarshal(b, &f)

	if err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println(f)
	}
	/*
		map[string]interface{} {
			"Name": "Wednesday",
			"Age":  6,
			"Parents": []interface{} {
				"Gomez",
				"Morticia",
			},
		}
	*/

	// 访问数据, 要是用断言
	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	var fm FamilMember
	err2 := json.Unmarshal(b, &fm)
	if err2 != nil {
		fmt.Println("error2")
	} else {
		fmt.Println("Name is: ", fm.Name)
		fmt.Println("Age is: ", fm.Age)
		fmt.Println("Parents are: ", fm.Parents)
	}
}
