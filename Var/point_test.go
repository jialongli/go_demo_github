package Var

import (
	"fmt"
	"testing"
)

func TestPoin2t(t *testing.T) {
	a := []int{1, 2}
	b := a
	fmt.Printf("a %p\n", a)
	fmt.Printf("&a %p\n", &a)
	fmt.Printf("==============\n")

	fmt.Printf("b %p\n", b)
	fmt.Printf("&b %p\n", &b)

	fmt.Printf("==============\n")

	c := make(chan string)
	d := c
	fmt.Printf("c %p\n", c)
	fmt.Printf("&c %p\n", &c)

	fmt.Printf("d %p\n", d)
	fmt.Printf("&d %p\n", &d)
	fmt.Printf("==============\n")

	m := 5
	n := m

	fmt.Printf("m %p\n", m)
	fmt.Printf("&m %p\n", &m)

	fmt.Printf("n %p\n", &n)
	fmt.Printf("&n %p\n", &n)

}
