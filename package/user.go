package mypackage

import (
	"fmt" //在$GOROOT/src目录下有fmt包
)

func init() {
	fmt.Println("enter package/user")
}

type User struct {
	Name string
	Age  int
	Sex  byte
}

func (u *User) Say() {
	fmt.Printf("Hi, I'm %s and %d yesrs old\n", u.Name, u.Age)
}
