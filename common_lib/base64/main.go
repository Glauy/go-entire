package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func base_64() {
	bs := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	str := base64.StdEncoding.EncodeToString(bs)
	fmt.Println(str)
	if cont, err := base64.StdEncoding.DecodeString(str); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cont)
	}

	//经常把小图片序列化成base64编码
	if fin, err := os.Open("data/gopher.png"); err != nil {
		fmt.Println(err)
	} else {
		defer fin.Close()
		bs := make([]byte, 10*1024)
		n, _ := fin.Read(bs)
		str := base64.StdEncoding.EncodeToString(bs[:n])
		fmt.Println(str)
	}
}

func main() {
	base_64()
}
