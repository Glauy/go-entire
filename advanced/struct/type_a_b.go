package main

import "fmt"

/**
类型转换
// 转换为Fish
td := Fish(fake)
// 真的变成了鱼
td.Swim()
*/
func main() {
	fake := FakeFish{}
	// fake 无法调用原来 Fish 的方法
	// 这一句会编译错误
	//fake.Swim()
	fake.FakeSwim()

	// 转换为Fish
	td := Fish(fake)
	// 真的变成了鱼
	td.Swim()

	sFake := StrongFakeFish{}
	// 这里就是调用了自己的方法
	sFake.Swim()

	td = Fish(sFake)
	// 真的变成了鱼
	td.Swim()
}

// FakeFish 定义了一个新类型，注意是新类型
type FakeFish Fish

func (f FakeFish) FakeSwim() {
	fmt.Printf("我是山寨鱼，嘎嘎嘎\n")
}

// StrongFakeFish 定义了一个新类型
type StrongFakeFish Fish

func (f StrongFakeFish) Swim() {
	fmt.Printf("我是华强北山寨鱼，嘎嘎嘎\n")
}

type Fish struct {
}

func (f Fish) Swim() {
	fmt.Printf("我是鱼，在游泳\n")
}
