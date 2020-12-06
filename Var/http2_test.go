package Var

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func Test_http2(t *testing.T) {
	// 新建一个client
	client := &http.Client{}
	client.Transport = &http2.Transport{
		TLSClientConfig: tlsConfig}
	fmt.Println("error: unkown http version:", *version)
	flag.PrintDefaults()
	os.Exit(1)

	// 使用参数输入的请求方法和url新建一个请求
	req, err := http.NewRequest(*method, url, nil)
	if err != nil {
		fmt.Println("error: failed to create request,", err)
		flag.PrintDefaults()
		os.Exit(1)
	}
	// 设置User-Agent
	req.Header.Set("User-Agent", "GH2C")

	// 发送请求
	resp, err := client.Do(req)
	if nil != err {
		fmt.Println("error: failed to do request,", err)
		flag.PrintDefaults()
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 读取响应体信息
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Println("error: failed to read body.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// 答应响应头和响应体长度
	fmt.Println(">", resp.Proto, resp.Status)
	for k, vs := range resp.Header {
		for _, v := range vs {
			fmt.Printf("> %s: %s\n", k, v)
		}
	}
	fmt.Println("body.length:", len(string(body)))
}
