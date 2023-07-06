package main

import (
	"encoding/json"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  byte `json:"gender"`
}

// 反射
func main() {

	user := User{
		Name: "钱钟书",
		Age:  57,
		Sex:  1,
	}

	json.Marshal(user) //返回 {"Name":"钱钟书","Age":57,"gender":1}

	reflect.Type()
}


