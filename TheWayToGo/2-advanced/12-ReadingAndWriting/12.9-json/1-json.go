package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Address   []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	// JSON format
	js, _ := json.Marshal(vc)
	fmt.Printf("Json format: %s\n", js)
	err := os.WriteFile("vcard.json", js, 0666)
	if err == nil {
		fmt.Println("json文件写入成功")
	} else {
		fmt.Println("json文件写入失败")
	}
	// using an encoder
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err1 := enc.Encode(vc)
	if err1 != nil {
		log.Println("Error in encoding json")
	}
	fmt.Println(err1)
}
