package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Addresses struct {
	Type    string
	City    string
	Country string
}

type VCards struct {
	FirstName string
	LastName  string
	Addresses []*Addresses
	Remark    string
}

func main() {
	pa := &Addresses{"private", "Aartselaar", "Belgium"}
	wa := &Addresses{"work", "Boom", "Belgium"}
	vc := VCards{"Jan", "Kersschot", []*Addresses{pa, wa}, "none"}

	js, _ := json.Marshal(vc)

	err := os.WriteFile("vcard.json", js, 0666)
	if err == nil {
		fmt.Println("json文件写入成功")
	} else {
		fmt.Println("json文件写入失败")
	}
}
