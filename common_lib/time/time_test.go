package main

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	time.Now().Format("2006-01-02 15:04:05")
}
