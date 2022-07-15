package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	stu := struct {
		Name string
		Age  int
		Sex  string
	}{
		Name: "cold bin",
		Age:  20,
		Sex:  "man",
	}
	//捕获动态类型
	t := reflect.TypeOf(stu)
	//校验动态类型
	if kind := t.Kind(); kind != reflect.Struct {
		log.Fatalf("This program expects to work on a struct; we got a %v instead.", kind)
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("Field %03d: %-10.10s %v\n", i, f.Name, f.Type.Kind())
	}
}
