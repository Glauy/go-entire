package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func compress() {
	// source file
	sFile, err := os.Open("data/readme.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	stat, _ := sFile.Stat()
	fmt.Printf("压缩前文件大小 %dB\n", stat.Size())

	// dest file
	dFile, err := os.OpenFile("readme.zlib", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, 1024)
	writer := zlib.NewWriter(dFile) //压缩写入
	for {
		n, err := sFile.Read(bs)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
			}
		} else {
			writer.Write(bs[:n])
		}
	}
	writer.Close()
	dFile.Close()
	sFile.Close()

	cFile, err := os.Open("readme.zlib")
	if err != nil {
		fmt.Println(err)
		return
	}
	stat, _ = cFile.Stat()
	fmt.Printf("压缩后文件大小 %dB\n", stat.Size())

	reader, err := zlib.NewReader(cFile) //解压
	io.Copy(os.Stdout, reader)           //把一个流拷贝到另外一个流
	reader.Close()
	cFile.Close()
}
