package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		skipParam string
		name      string
		skip      int
		err       error
	)
	//tips: strconv.Atoi 字符串转换成为整数
	if skip, err = strconv.Atoi(skipParam); err != nil {
		skip = 0
	}

	fmt.Println(name, skip)
}
