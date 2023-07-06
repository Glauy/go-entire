package main

import (
	"fmt"
	"log"
	"os"
)

//打日志
func logger() {
	log.Printf("%d+%d=%d\n", 3, 4, 3+4)
	log.Println("Hello Golang")
	// log.Fatalln("Bye, the world") //日志输出后会执行os.Exit(1)

	//以append方式打开日志文件
	fout, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open log file failed: %v\n", err)
	}
	defer fout.Close()
	logWriter := log.New(fout, "[MY_BIZ]", log.Ldate|log.Lmicroseconds) //通过flag参数定义日志的格式，时间精确到微秒1E-6s
	logWriter.Printf("%d+%d=%d\n", 3, 4, 3+4)
	logWriter.Println("Hello Golang")
	// logWriter.Fatalln("Bye, the world")
}

func file() error {
	f, err := os.Open("test.log")
	if err != nil {
		return err
	}
	defer f.Close()
	stat, _ := f.Stat()
	fmt.Printf("文件大小 %dB\n", stat.Size())
	return nil
}

func main() {
	//logger()
	file()
}
