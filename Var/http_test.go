package Var

import (
	"os"
	"strconv"
	"testing"
)

func TestHttpRouter(t *testing.T) {
	mode := 777
	um, _ := strconv.ParseInt(strconv.Itoa(mode), 8, 0)

	os.Mkdir("/tmp/ljltest2", os.FileMode(um))
}
