package main

import "fmt"

func main() {

	var arr4 = [...]int{
		99: 39, // 将第100个元素(下标值为99)的值赋值为39，其余元素值均为0
	}
	fmt.Printf("%T\n", arr4) // [100]int
}
