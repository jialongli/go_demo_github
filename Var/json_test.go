package Var

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

//func TestT2(t *testing.T) {
//	s := `{"code":"200"}`
//	b := []byte(s)
//	m := &CcmdbResponse{}
//	re := json.Unmarshal(b, &m)
//	t.Log(m, re)
//}

type CcmdbResponse struct {
	Code string `json:"code"`
}

type info2 struct {
	ip  string `json:"ip"`
	idc string `json:"idc"`
}

func TestStringToMap(t *testing.T) {
	headers := make(map[string]string)
	headers["options"] = `{"version":"1.2.1.1","tag":"order"}`
	//获取header中的options
	options := make(map[string]string)
	if optionsStr, ok := headers["options"]; ok {
		//TODO 这为什么要用&
		error := json.Unmarshal([]byte(optionsStr), &options)
		if error != nil {
			t.Error()
		}
	}

}

type Environment struct {
	EnvName     string `json:"envName"`
	Psm         string `json:"psm"`
	DeployType  string `json:"deployType,default"`
	EnvType     string `json:"envType",default:"TEST"`
	Domain      string `json:"domain"`
	ContextPath string `json:"contextPath"`
	InteType    string `json:"inteType",default:"TLB"`
}

func TestJson(t *testing.T) {
	j := `{
    "envName":"test222",
    "psm":"ddd",
    "domain":"domaintt",
    "contextPath":"sdfsdfa"
}`
	en := Environment{}
	json.Unmarshal([]byte(j), &en)
	t.Log(en)
}

func TestHttp(t *testing.T) {

	url := `http://10.100.201.103:8001/dsf/tcbase/dsf/demo/dotnet/service1/neo`
	// 网络数据传输， 数据复制完成后，需要释放连接和对象资源
	remoteResp := fasthttp.AcquireResponse()
	req := fasthttp.AcquireRequest()
	h := &req.Header
	md.ConvertToHttpRequestHeader(h)
	h.Del("host")
	req.SetRequestURI(url)
	oldReqBody := req.SwapBody(body)
	defer func() {
		// 替换原缓冲区
		req.SwapBody(oldReqBody)
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(remoteResp)
	}()
	err := transport.FastHttp1TransportWithoutRelease(req, remoteResp, rule.DeadLineTime())
}
