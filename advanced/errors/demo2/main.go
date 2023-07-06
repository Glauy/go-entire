package main

import (
	"errors"
	"fmt"
)

func New(text string) error {
	return errorString{text}
}

// errorString is a trivial implementation of errors.
type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}
func main() {
	if New("EOF") == New("EOF") {
		fmt.Println("自定义errorString，使用值类型")
	}

	if errors.New("EOF") == errors.New("EOF") {
		fmt.Println("内置errorString，使用引用类型")
	}
}
