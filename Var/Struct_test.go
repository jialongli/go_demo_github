package Var

import (
	"fmt"
	"strconv"
	"testing"
)

/**
duck type: 只要长的像鸭子,就是鸭子类型

不需要显示声明继承接口,只要实现了这个接口的方法,(方法签名)
*/
type sayHello interface {
	sayHello()
}

type user struct {
	name string
	age  int
}

func (user *user) sayHello() {
	fmt.Print("hello im" + user.name + " my age is" + strconv.Itoa(user.age))
}
func TestStruct(t *testing.T) {
	user1 := new(user)
	user1.age = 12
	user1.name = "ljl"
	user1.sayHello()
}
