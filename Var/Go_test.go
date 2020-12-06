package Var

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	for i := 0; i < 10; i++ {
		go goFunc(i, 10)
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("不安全啊" + strconv.Itoa(i))
		}(i)
	}

	time.Sleep(2000)
}

func goFunc(i int, j int) {
	fmt.Println(i, j)
}

func s() {
	tr := &http2.Transport{AllowHTTP: true,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		}}
	cli := http.Client{Transport: tr}
	resp, err := cli.Get("http://localhost:5678/")
	log.Println(err, resp)
}
