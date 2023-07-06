package main

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	var fileName = "x.txt"
	newFile, err := os.Create("/tmp/" + fileName)
	//io.Copy()
	newFile.Seek(0, 0)
}
