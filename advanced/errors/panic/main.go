package main

import (
	"fmt"
)

func div(a, b int) {
	if b < 0 {
		panic("除数需要大于0")
	}
	fmt.Println("余数为：", a/b)

}

func div2(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到异常：%s\n", r)
		}
	}()
	fmt.Println("余数为：", a/b)
}

func div3(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到异常：%s\n", r)
		}
	}()

	if b < 0 {
		panic("除数需要大于0")
	}
	fmt.Println("余数为：", a/b)
}

func main() {
	// 捕捉主动的异常
	//div(10, -1)
	//time.Sleep(10 * time.Second)
	//div2(10, 0)
	div3(10, -1)
}
