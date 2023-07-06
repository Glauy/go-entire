package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) m1() {
	fmt.Printf("enter m1：address is %x\n", unsafe.Pointer(&s.Name))
}

func (s *Student) m2() {
	fmt.Printf("enter m2：address is %x\n", unsafe.Pointer(&s.Name))
}

type errorString struct {
	s string
}

// 错误处理章节中 errorString 返回的是指针，避免了将两个错误识别为一样的
func test() {
	e := errorString{"EOF"}
	e2 := errorString{"EOF"}
	if e == e2 {
		fmt.Println("xx")
	}

	p := &errorString{"EOF"}
	p2 := &errorString{"EOF"}
	if p == p2 {
		fmt.Println("yy")
	}
}

func main() {
	test()
	//student := Student{"zs", 23}
	//fmt.Printf("main: address is %x\n", unsafe.Pointer(&student.Name))
	//student.m1()
	//student.m2()
}
