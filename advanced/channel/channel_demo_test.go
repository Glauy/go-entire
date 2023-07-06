package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			fmt.Println(v)
		}
	}
}
