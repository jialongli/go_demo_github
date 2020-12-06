package Var

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

	c := server1()
	server2()
	time.Sleep(3000)
	fmt.Println("我开始读取channel了" + <-c)
	//fmt.Println("我再次读取channel了" + <-c)
	fmt.Println("我再次读取channel结束了")
}

func server1() chan string {
	//这种为什么错误?因为我在这个方法里创建了一个channel,但是这个channel在这个方法里始终没有读取操作,就会导致执行该方法的协程一直阻塞在 c<-"wan"这里
	//很可怕,所以启动报错了
	//===========fatal error: all goroutines are asleep - deadlock!==================
	c := make(chan string)
	fmt.Println("server1 我把wan写入了chan")
	c <- "wan"
	fmt.Println("server1已经结束了")
	return c
}

func server2() {
	fmt.Println("我是server2")
}
