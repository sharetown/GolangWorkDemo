package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Addr string `json:"addr,omitempty"`
}

func main() {
	p1 := Person{
		Name: "Jack",
		Age:  22,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// p1 没有 Addr，就不会打印了
	fmt.Printf("%s\n", data1)

	// ================

	p2 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有
	fmt.Printf("%s\n", data2)

	p3 := Person{}

	data3, err := json.Marshal(p3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data3)

	// 三种获取 field
	field1,_ := reflect.TypeOf(p1).FieldByName("Name")
	// field2 := reflect.ValueOf(p2).Type().Field(2)         // i 表示第几个字段
	// field3 := reflect.ValueOf(&p3).Elem().Type().Field(3) // i 表示第几个字段

	// 获取 Tag
	tag1 := field1.Tag

	// 获取键值对
	labelValue1 := tag1.Get("json")
	labelValue2, _ := tag1.Lookup("json")

    fmt.Printf("%s\n",labelValue1)
    fmt.Printf("%s\n",labelValue2)


}
