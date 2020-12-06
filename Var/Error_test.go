package Var

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	//这里返回了一个error的实例,因为error是接口,所以只要定义一个struct,且这个struct有一个方法,实现了Error,就ok了
	//errors.New("hhhh")

	value, result := valid(00)
	if result == nil {
		fmt.Print(value)
	} else {
		fmt.Println("我错了")
	}

}

func valid(t int) (int, error) {
	if t > 100 {
		return 100, errors.New("错误的值,>100")
	} else {
		return t, nil
	}
}

var LARGER_ERROR = errors.New("lager woccccc")
var SMALL_ERROR = errors.New("small woccccc")

func validWithDetailError(t int) (int, error) {
	if t > 100 {
		return t, LARGER_ERROR
	} else if t < 0 {
		return t, SMALL_ERROR
	} else {
		return t, nil
	}
}

func TestHH(t *testing.T) {
	var error error = nil
	s, error := test1()
	t.Log(s, error)
}

func test1() (string, error) {
	return "sss", errors.New("hoho")
}
