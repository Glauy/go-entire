package main

import "fmt"

func add4int(a, b int) int {
	return a + b
}
func add4float32(a, b float32) float32 {
	return a + b
}
func add4string(a, b string) string {
	return a + b
}

type Addable interface{
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128,string
}

func add[T Addable](a,b T) T {
	return a + b
}

func main() {
	var a, b int = 2, 5
	c, d := "3214", "few"
	fmt.Println(add4int(a, b))
	fmt.Println(add4string(c, d))

	//fmt.Println(add(a, b))
	//fmt.Println(add(c, d))
}
