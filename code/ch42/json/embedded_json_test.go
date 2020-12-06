package jsontest

import (
	"reflect"
	"testing"
)

var jsonStr = `{
	"basic_info":{
	  	"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}	`

func TestEmbeddedJson(t *testing.T) {
	t.Log(reflect.TypeOf(jsonStr))

	//e := new(Employee)
	//err := json.Unmarshal([]byte(jsonStr), e)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println(*e)
	//if v, err := json.Marshal(e); err == nil {
	//	fmt.Println(string(v))
	//} else {
	//	t.Error(err)
	//}

}
