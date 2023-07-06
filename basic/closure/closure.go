package main

import (
	"fmt"
	"time"
)

func main() {

	i := 13
	a := func() {
		fmt.Printf("i is %d \n", i)
	}
	a()

	fmt.Println(ReturnClosure("Tom")())

	Delay()
	time.Sleep(time.Second)
}

// ReturnClosure 这种写法很常见
func ReturnClosure(name string) func() string {
	return func() string {
		return "Hello, " + name
	}
}

// Delay 闭包存在延迟绑定问题
func Delay() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			fmt.Printf("hello, this is : %d \n", i)
		})
	}

	for _, fn := range fns {
		fn()
	}
}
