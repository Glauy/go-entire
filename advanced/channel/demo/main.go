package main

import "fmt"

func demo1() {
	//nil channel在select里面能正常运行，拿不到数据
	var c1, c2 chan int // c1 and c2 = nil
	//c1 := make(chan int)
	//c2 := make(chan int)

	select {
	case n := <-c1:
		fmt.Println("received from c1:", n)
	case n := <-c2:
		fmt.Println("received from c2:", n)
	default:
		fmt.Println("no value received")
	}
}

func main() {
	demo1()

}
