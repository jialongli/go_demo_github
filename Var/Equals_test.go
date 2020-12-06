package Var

import (
	"testing"
)

func TestEquals(t *testing.T) {
	a := [...]int{1, 2}
	b := [...]int{1, 2}
	//c:= [...]int{1,2,3}
	d := [...]int{1, 3}
	t.Log(a == b)
	t.Log(a == d)

}

func TestCase(t *testing.T) {
	//for的条件 不要加() 否则报错
	//
	for i := 0; i < 10; i++ {
		switch i {
		case 1:
			t.Log("1")
		case 2:
			t.Log("2")
		case i == 3:
			t.Log("i==3")

		}
	}
}

func TestCase2(t *testing.T) {
	//for的条件 不要加() 否则报错
	//
	for i := 0; i < 10; i++ {
		switch {
		case i == 1:
			t.Log("1")
		case i == 2:
			t.Log("2")
		case i == 3:
			t.Log("i==3")

		}
	}
}
