package Var

import (
	"fmt"
	"testing"
	"unsafe"
)

type Pet struct {
}

func (pet *Pet) speak() {
	fmt.Print("pet speak")
}

func (pet *Pet) speakTo(name string) {
	fmt.Print("pet speakto  " + name)

}

type Dog struct {
	pet *Pet
}

func (dog *Dog) speak() {
	dog.pet.speak()
}

func (dog *Dog) speakTo(name string) {
	dog.pet.speakTo(name)
}

func TestP(t *testing.T) {
	dog := new(Dog)
	dog.speakTo("ljl")
}

func TestObject(t *testing.T) {
	d1 := new(Dog)
	d2 := Dog{}
	d3 := &d2
	// unsafe.Pointer需要传入一个指针,所以
	fmt.Printf("%X\n", unsafe.Pointer(d1))
	fmt.Printf("%X\n", unsafe.Pointer(d2))
	fmt.Printf("%X\n", unsafe.Pointer(d3))

}
