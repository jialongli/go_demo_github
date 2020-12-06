package Var

import "testing"

const (
	Sunday = 1 + iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
)

func TestTT(t *testing.T) {
	t.Log(Sunday, Tuesday)
}
