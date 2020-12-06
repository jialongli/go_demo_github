package Var

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]string{"one": "1", "two": "2"}
	s, ok := m["3"]
	if ok {

	}
	t.Log(s)

	if value, exit := m["3"]; exit {
		t.Log(value)
	}

	for key, value := range m {
		t.Log(key, value)
	}
}

//map中可以存放方法,那么可以实现一个工厂模式
func TestFactory(t *testing.T) {
	funcMap := map[string]func(op int) int{}
	funcMap["1"] = func(op int) int {
		return op
	}
	funcMap["2"] = func(op int) int {
		return op * op
	}
	t.Log(funcMap["1"](2))
	t.Log(funcMap["2"](2))

}

func TestMail(t *testing.T) {
	//m := make(map[string]interface{})
	m := make(map[string]map[string]interface{})
	mail := `{"data":{"status":0,"message":"操作成功","data":"a0142cca43b94345ac58c8a04bca4c1f"}}`
	json.Unmarshal([]byte(mail), &m)
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	t.Log(m)
}
