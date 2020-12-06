package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	var unixConn net.Conn
	var err error
	unixConn, err = net.DialTimeout("unix", "/data/logs/uds/test_sock", 1*time.Second)
	if err != nil {
		fmt.Println("client error" + err.Error())
	}
	defer unixConn.Close()

	uc := unixConn.(*net.UnixConn)
	buf := make([]byte, 3)
	n, _ := uc.Read(buf)
	fmt.Println("收到返回,字节长度", strconv.Itoa(n)+",内容:"+string(buf))
}
