package main

import (
	"fmt"
	"time"
)

//版本一：for 通道关闭，接收方接收到通道零值
//版本二：for range
func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
}

// 定义channel，限定了缓冲区大小，并且发送缓冲区数据正好小于缓冲区大小，那么不定义接收者不会报错
func bufferedChannel() {
	c := make(chan int, 3)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	//c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func t() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
