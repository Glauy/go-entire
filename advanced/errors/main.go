package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) (err error) {
	file, err := os.OpenFile(path, os.O_CREATE, 0666)

	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}
	}()

	err = errors.New("this is a custom errors")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			// panic 的是err，pathError不能断言成功，为nil
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Error())
		}
		return
	}

	_, err = io.WriteString(file, "hello golang")
	return
}

func main() {
	err := WriteFile("test.txt")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
