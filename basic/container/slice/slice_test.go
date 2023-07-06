package main

import (
	"fmt"
	"testing"
	"unsafe"
)

//传入切片
//参数传递都是值传递，切片也不例外，只不过存在底层数组引用，所以方法内部修改切片底层数组值，外面也能看到！
func change(s []string) {
	fmt.Printf("lang: address is %x\n", unsafe.Pointer(&s))
	s[1] = "c++"
}

//测试切片传参，修改底层数组
func TestSliceParam(t *testing.T) {
	lang := []string{"golang", "java", "php", "js"}
	fmt.Printf("lang: address is %x\n", unsafe.Pointer(&lang))
	fmt.Printf("")
	fmt.Println(lang[1]) //java
	change(lang)
	fmt.Println(lang[1]) //c++
}

func TestNilSlice(t *testing.T) {
	var s []int
	fmt.Println(len(s), cap(s))
	fmt.Printf("s: address is %x\n", unsafe.Pointer(&s))
	s = append(s, 11)
	fmt.Printf("s: address is %x\n", unsafe.Pointer(&s))
	fmt.Println(len(s), cap(s))
}

func TestEmptySlice(t *testing.T) {

}

//测试扩容 2倍扩容
func TestSliceScale(t *testing.T) {
	var s []int
	s = append(s, 11)
	fmt.Println(len(s), cap(s)) //1 1
	s = append(s, 12)
	fmt.Println(len(s), cap(s)) //2 2
	s = append(s, 13)
	fmt.Println(len(s), cap(s)) //3 4
	s = append(s, 14)
	fmt.Println(len(s), cap(s)) //4 4
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) //5 8
}

//func main() {
//	// 示例1。
//	s1 := make([]int, 5)
//	fmt.Printf("The length of s1: %d\n", len(s1))
//	fmt.Printf("The capacity of s1: %d\n", cap(s1))
//	fmt.Printf("The value of s1: %d\n", s1)
//	s2 := make([]int, 5, 8)
//	fmt.Printf("The length of s2: %d\n", len(s2))
//	fmt.Printf("The capacity of s2: %d\n", cap(s2))
//	fmt.Printf("The value of s2: %d\n", s2)
//
//	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
//	s4 := s3[3:6]
//	fmt.Printf("The length of s4: %d\n", len(s4))
//	fmt.Printf("The capacity of s4: %d\n", cap(s4))
//	fmt.Printf("The value of s4: %d\n", s4)
//}
