package main

import (
	"fmt"
	"me-go/inteface/mock"
	"me-go/inteface/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "this mock content"}
	fmt.Printf("%T %v\n", r, r)

	r = real.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	fmt.Printf("%T %v\n", r, r)
}
