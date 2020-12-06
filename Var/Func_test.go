package Var

import "testing"

//map中可以存放方法,那么可以实现一个工厂模式
func TestFunc(t *testing.T) {
	testMany(t, 1, 2, 3)
}

func testMany(t *testing.T, opt ...int) int {
	for s := range opt {
		t.Log(s)
	}
	return 1
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("defer.....")
	}()
	t.Log("firset")
	panic("err")
}
