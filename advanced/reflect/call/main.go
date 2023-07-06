package main

import (
	"fmt"
	"reflect"
)

func Add(a int, b int) int {
	return a + b
}

func Add2(a int, b int) int {
	return a + b
}

func CallAddByReflect(f func(a int, b int) int) {
	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		return
	}

	argv := make([]reflect.Value, 2)
	argv[0] = reflect.ValueOf(1)
	argv[1] = reflect.ValueOf(2)

	result := v.Call(argv)
	fmt.Println(result[0])
}

func main() {
	CallAddByReflect(Add)

}
