package main

import "fmt"

func main() {

	var arr2 = [6]int{
		11, 12, 13, 14, 15, 16,
	} // [11 12 13 14 15 16]

	fmt.Println(arr2)
	var arr3 = [...]int{
		21, 22, 23,
	} // [21 22 23]
	fmt.Printf("%T\n", arr3) // [3]int
}
