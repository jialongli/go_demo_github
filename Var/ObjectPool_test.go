package Var

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type cat struct {
	name string
	age  int
}

type objectPool struct {
	catChan chan *cat
}

var p = new(objectPool)

/**
初始化池子以及几只猫
*/
func init() {
	p.catChan = make(chan *cat, 2)
	p.catChan <- &cat{}
	p.catChan <- &cat{}
}

func getCat() (*cat, error) {
	select {
	case cat := <-p.catChan:
		fmt.Println("我获取到了一只小猫咪")
		return cat, nil
	case <-time.After(time.Second * 1):
		return nil, errors.New("timeout")
	}
}

func release(cat *cat) error {
	select {
	case p.catChan <- cat:
		fmt.Println("我还回去了一只小猫咪")
		return nil
	default:
		fmt.Println("!!!!!!我还不回去了")
		return errors.New("")
	}
}

func TestFa(t *testing.T) {
	cat, _ := getCat()
	fmt.Println("get" + cat.name)
	cat1, _ := getCat()
	fmt.Println("get" + cat1.name)
	release(cat1)

	cat2, _ := getCat()
	fmt.Println("get" + cat2.name)
	fmt.Println("get" + cat.name)
}
