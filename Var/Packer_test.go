package Var

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Print("init hahah1")
}

func init() {
	fmt.Print("init hahah2")
}

func init() {
	fmt.Print("init hahah3")
}

func TestInit(t *testing.T) {
	t.Log("testIng")
}
