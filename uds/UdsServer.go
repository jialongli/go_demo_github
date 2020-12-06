package main

import (
	"fmt"
	"net"
	"syscall"
)

var (
	Sock_addr = "/data/logs/uds/test_sock"
)

func main() {
	syscall.Unlink(Sock_addr)

	//=====1.mosn启动后,来监听reconfig这个uds
	l, err := net.Listen("unix", Sock_addr)
	//注意这里,目录需要提前建好,但是不要建立文件,否则会报错,address has bind.哦 ,所以mosn里面每次要unlink一下
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer l.Close()

	ul := l.(*net.UnixListener)
	for {
		//=======2.当accept到连接后.
		uc, err := ul.AcceptUnix()
		if err != nil {
			return
		}
		//3.写入一个字符串数据作为回应
		_, err = uc.Write([]byte("ljl"))
		if err != nil {
			continue
		}
		//4.可以关闭连接了
		uc.Close()
	}
}
