package Var

import (
	"math"
	"testing"
)

/**
类型的别名
*/
type aliseInt int64

func TestType(t *testing.T) {
	a := 3
	var b int32 = int32(a)
	t.Log(b)
	t.Log(math.MaxFloat32)
}

func TestPoint(t *testing.T) {
	a := 32
	b := &a
	*b = *b + 1
	t.Log(a, b)
}

var container = []string{"zero", "one", "two"}

func TestType2(t *testing.T) {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	v, ok := interface{}(container).(string)
}
