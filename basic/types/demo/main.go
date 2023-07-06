package main

import (
	"fmt"
	"unsafe"
)

//类型 字节长度
func demo1() {
	i := int32(0)
	//i := int(0)
	p := &i
	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(unsafe.Sizeof(p))
}

func demo2() {
	fmt.Println(unsafe.Sizeof("慕课网"))
	fmt.Println(unsafe.Sizeof("慕课网moody老师"))
}

func demo3() {
	var a *int
	fmt.Println(a == nil)
}
func main() {
	//demo2()
	demo3()

}
