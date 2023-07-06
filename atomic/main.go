package main

import "sync/atomic"

func main() {
	var c int32 = 10
	atomic.CompareAndSwapInt32(&c, 10, 100)

	//相当于
	if c == 10 {
		c = 100
	}

}
