package testSidecar

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestCons(t *testing.T) {
	//a := uint32(16 << 20)
	//t.Log(a)
	//
	//b := make([]int, 2)
	//c := b[:0]
	//t.Log(c)
	//
	//l := new(map[string]ss)
	//fmt.Println(l)
	//
	//l1 := make(map[string]struct{ name string })
	//fmt.Println(l1)
	//
	//
	//m := make(map[string]chan ss)
	//fmt.Println(m)

	fmt.Println(time.Now().Clock())

	m2 := make(map[string]string, 2)
	fmt.Println(len(m2))
	for one := range m2 {
		fmt.Println("fdsafas")
		fmt.Println(one)
	}

	str := "http://helllo"
	str = strings.Replace(str, "http://", "", -1)
	fmt.Println(str)

}

type ss struct {
	name string
}
