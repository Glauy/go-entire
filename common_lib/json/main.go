package main

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
)

type Student struct {
	Name string
	Age  int32
}

func main() {
	stu := Student{Name: "sz", Age: 25}
	bs, _ := json.Marshal(stu)
	fmt.Println(string(bs))

	bs, _ = sonic.Marshal(stu)
	fmt.Println(string(bs))
}
