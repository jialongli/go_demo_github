package Var

import (
	"fmt"
	"testing"
)

func TestChannel2(t *testing.T) {
	c := make(chan string)
	fmt.Printf("创建时候的返回的地址,应该是引用地址%p\n", c)
	fmt.Printf("创建的实际地址,应该和方法里的一样%p\n", &c)
	s1(&c)
	//s2()
	//time.Sleep(3000)
	//fmt.Println("我开始读取channel了" + <-c)
	////fmt.Println("我再次读取channel了" + <-c)
	//fmt.Println("我再次读取channel结束了")
}

func s1(c *chan string) {
	b := c
	fmt.Printf("'%p\n", b)
	fmt.Printf("通道的实际存储地址 %p\n", &b)
	fmt.Printf("方法里存储的引用地址%p\n", c)
	fmt.Printf("通道的实际存储地址 %p\n", &c)

	//fmt.Println("server1 我把wan写入了chan")
	//*c <- "wan"
	//fmt.Println("server1已经结束了")
}

func s2() {
	fmt.Println("我是server2")
}
