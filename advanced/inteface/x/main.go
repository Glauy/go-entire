package main

type Interface1 interface {
	M1()
}
type Interface2 interface {
	M1(string)
	M2()
}

type Interface3 interface {
	Interface1
	Interface2 // 编译器报错：duplicate method M1
	M3()
}

func main() {

}
