package main

import (
	"fmt"
	"time"
)

//其实，我们只需要将 field 类型 print 方法的 receiver 类型由 *field 改为 field 就可以了。
type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print()
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go v.print()
	}

	time.Sleep(3 * time.Second)
}
