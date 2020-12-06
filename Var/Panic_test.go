package Var

import (
	"testing"
)

/**
osExit退出不会
*/
func TestOsExit(t *testing.T) {
	defer func() {
		t.Log("我是defer")
	}()
	//os.Exit(-1)
}

/**
osExit退出不会
*/
func TestPanic(t *testing.T) {
	defer func() {
		t.Log("我是defer")
	}()
	panic("hahf")
}
