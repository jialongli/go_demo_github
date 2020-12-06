package Var

import (
	"encoding/json"
	"strings"
	"testing"
)

func Test_s(t *testing.T) {
	s := "client++ser+api+ins"
	i := strings.Index(s, "++")
	key := s[i+2:]
	t.Log(key)
}

func Test_xx(t *testing.T) {
	getUri("/name/hello", `{"key1":"v1","key2":"v2"}`)
}

/**
获取实际url.
*/
func getUri(apiMap string, param string) interface{} {
	paramStr := ""
	if param != "" {
		paramStr = paramStr + "?"
		paramMap := make(map[string]string)
		json.Unmarshal([]byte(param), &paramMap)
		for k, v := range paramMap {
			paramStr = paramStr + k + "=" + v + "&"
		}
		paramStr = paramStr[:len(paramStr)-1]
	}
	return apiMap + paramStr
}
