package Var

import (
	"testing"
	"time"
)

func TestCh2(t *testing.T) {
	c := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c <- "dfasdfas"
	}()
	t.Log("我要开始选择了")
	select {
	case <-c:
		t.Log("c来消息了哈哈哈")
		//哈哈,典型的错误,第二次获取通道里的值,就会阻塞在这里
		//fmt.Println("c来消息了哈 <-c" + <-c)
	case <-time.After(time.Second * 1):
		t.Log("1s后我执行了")
	}

	t.Log("结束")

	time.Sleep(time.Second * 5)
}
