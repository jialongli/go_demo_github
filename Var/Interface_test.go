package Var

import (
	"fmt"
	"testing"
)

type speak interface {
	speak(s string)
}

type duck struct {
}

func (d *duck) speak(s string) {
	fmt.Print("鸭子 speak")
}

type dog struct {
}

func (d *dog) speak(s string) {
	fmt.Println("狗 speak")
}

func speakTest(i speak) {
	i.speak("xxx")
}

func TestSpeak(t *testing.T) {
	dog1 := &dog{}
	dog := new(dog)
	duck := new(duck)
	speakTest(dog)
	speakTest(duck)
	speakTest(dog1)

}

//可以用空接口表示任何类型 类似java里 object 这个参数是一个具体的接口
func hh(i interface{ speak(s string) }) {

}

//可以用空接口表示任何类型 类似java里 object
func hhNoFunc(i interface{}) {
	if v, result := i.(int); result {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
}

func TestFF(t *testing.T) {
	hhNoFunc("13")
}
