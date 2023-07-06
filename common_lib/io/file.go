package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//格式化输出
func format() {
	var i int = 1234
	var f float32 = 3.1415
	var stu struct {
		Name string
		Age  int
	}
	stu.Name = "张三"
	stu.Age = 18
	fmt.Printf("%b\n", i) //二进制表示
	fmt.Printf("%d\n", i)
	fmt.Printf("%8d\n", i)  //左边补空格，补够8位
	fmt.Printf("%08d\n", i) //左边补0，补够8位
	fmt.Printf("%f\n", f)   //默认6位小数
	fmt.Printf("%.8f\n", f)
	fmt.Printf("%e\n", f) //科学计数法，默认6位小数
	fmt.Printf("%.8e\n", f)
	fmt.Printf("%g\n", f)     //根据实际情况采用 %e 或 %f 格式（获得更简洁、准确的输出）
	fmt.Printf("%t\n", 3 > 9) //true或false
	fmt.Printf("%s\n", stu.Name)

	fmt.Printf("%T\n", stu)  //输出类型，包括结构体各个字段的名称和类型。struct { Name string; Age int }
	fmt.Printf("%v\n", stu)  //输出结构体各个字段的值。{张三 18}
	fmt.Printf("%+v\n", stu) //还会带上字段名。{Name:张三 Age:18}
	fmt.Printf("%#v\n", stu) //差不多相当于先调用%T，再调%+v。struct { Name string; Age int }{Name:"张三", Age:18}
}

//从标准输入读入数据
func scan() {
	fmt.Println("please input a word")
	var word string
	fmt.Scan(&word) //读入第1个空格前的单词
	fmt.Println(word)

	fmt.Println("please input two word")
	var word1 string
	var word2 string
	fmt.Scan(&word1, &word2) //读入多个单词，空格分隔。如果输入了更多单词会被缓存起来，丢给下一次scan
	fmt.Println(word1, word2)

	fmt.Println("please input an int")
	var i int
	fmt.Scanf("%d", &i) //类似于Scan，转为特定格式的数据
	fmt.Println(i)
}

func read_file() {
	if fin, err := os.Open("data/digit.txt"); err != nil {
		fmt.Printf("open file faied: %v\n", err) //比如文件不存在
	} else {
		defer fin.Close() //别忘了关闭文件句柄

		//读二进制文件
		cont := make([]byte, 10)
		if n, err := fin.Read(cont); err != nil { //读出len(cont)个字节，返回成功读取的字节数
			fmt.Printf("read file failed: %v\n", err)
		} else {
			fmt.Println(string(cont[:n]))
			if m, err := fin.ReadAt(cont, int64(n)); err != nil { //从指定的位置开始读len(cont)个字节
				fmt.Printf("read file failed: %v\n", err)
			} else {
				fmt.Println(string(cont[:m]))
			}
			fin.Seek(int64(n), 0) //whence: 0从文件开头计算偏移量，1从当前位置计算偏移量，2到文件末尾的偏移量
			if n, err = fin.Read(cont); err != nil {
				fmt.Printf("read file failed: %v\n", err)
			} else {
				fmt.Println(string(cont[:n]))
			}
		}

		//读文本文件建议用bufio.Reader
		fin.Seek(0, 0) //定位到文件开头
		reader := bufio.NewReader(fin)
		for { //无限循环
			if line, err := reader.ReadString('\n'); err != nil { //指定分隔符
				if err == io.EOF {
					if len(line) > 0 { //如果最后一行没有换行符，则此时最后一行就存在line里
						fmt.Println(line)
					}
					break //已读到文件末尾
				} else {
					fmt.Printf("read file failed: %v\n", err)
				}
			} else {
				line = strings.TrimRight(line, "\n") //line里面是包含换行符的，需要去掉
				fmt.Println(line)
			}
		}
	}
}

func write_file() {
	//OpenFile()比Open()有更多的参数选项。os.O_WRONLY以只写的方式打开文件，os.O_TRUNC把文件之前的内容先清空掉，os.O_CREATE如果文件不存在则先创建，0666新建文件的权限设置
	if fout, err := os.OpenFile("data/verse.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
		fmt.Printf("open file faied: %v\n", err)
	} else {
		defer fout.Close() //别忘了关闭文件句柄

		//写文本文件建议使用
		writer := bufio.NewWriter(fout)
		writer.WriteString("明月多情应笑我")
		writer.WriteString("\n") //需要手动写入换行符
		writer.WriteString("笑我如今")
		writer.Flush() //buffer中的数据量积累到一定程度后才会真正写入磁盘。调用Flush强行把缓冲中的所有内容写入磁盘
	}
}
