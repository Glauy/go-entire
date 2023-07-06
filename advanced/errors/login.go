package main

import (
	"errors"
	"fmt"
)

// 这是一个自定义类型错误示例
type LoginError string

const (
	ErrInvalidCredentials LoginError = "Invalid credentials"
	ErrAccountLocked      LoginError = "Account locked"
	ErrAccountInactive    LoginError = "Account inactive"
)

func (le LoginError) Error() string {
	return string(le)
}

func Login(username, password string) error {
	// 模拟登录验证逻辑
	if username != "admin" || password != "password" {
		return errors.New(string(ErrInvalidCredentials))
	}

	// 检查账户状态等其他条件
	// ...

	return nil
}

func T() {
	err := Login("john", "pass")
	if err != nil {
		if loginErr, ok := err.(LoginError); ok {
			switch loginErr {
			case ErrInvalidCredentials:
				fmt.Println("Error:", loginErr)
				// 处理无效凭证的情况
			case ErrAccountLocked:
				fmt.Println("Error:", loginErr)
				// 处理账户被锁定的情况
			case ErrAccountInactive:
				fmt.Println("Error:", loginErr)
				// 处理账户不活跃的情况
			default:
				fmt.Println("Unknown error:", loginErr)
				// 处理其他未知错误
			}
		} else {
			fmt.Println("Error:", err)
			// 处理非自定义错误
		}
		return
	}

	// 登录成功
	fmt.Println("Login successful")
}
