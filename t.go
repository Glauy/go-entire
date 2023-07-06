package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

//0 2 1 1 2
func p1(sig chan struct{}) {
	for i := 1; i < math.MaxInt; i += 2 {
		fmt.Println("奇数：", i)
		sig <- struct{}{}
		time.Sleep(time.Second)
	}
}

func p2(sig chan struct{}) {
	for i := 2; i < math.MaxInt; i += 2 {
		<-sig
		fmt.Println("偶数：", i)

		time.Sleep(time.Second)
	}
}

func main() {
	sig := make(chan struct{})
	group := sync.WaitGroup{}
	group.Add(2)
	go p1(sig)
	go p2(sig)
	group.Wait()
}
