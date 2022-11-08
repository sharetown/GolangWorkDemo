package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num string = "ab"

	t := reflect.TypeOf(num)
	v := reflect.ValueOf(num)
    
    // reflect.TypeOf(num) 返回的是接口变量num的反射对象，其类型是*reflect.rtype，值是string
	fmt.Printf("%T\n", t)
	fmt.Printf("%v\n", t)

    // reflect.ValueOf(num)返回的是接口变量num的反射对象,其类型是reflect.Value，值是ab
	fmt.Printf("%T\n", v)
	fmt.Printf("%v\n", v)

    // v.Interface()将反射对象reflect.Value转换为接口变量，其类型是string，值是ab
	i:= v.Interface().(string)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", i)

}
