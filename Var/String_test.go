package Var

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s1 := "woaibeijign"
	t.Log(s1)
	r := strings.Split(s1, "a")
	t.Log(r)
}
