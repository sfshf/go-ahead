package singleton

import "testing"

func TestSingleton(t *testing.T) {
    ins1 := Singleton()
    ins2 := Singleton()
    if ins1 != ins2 {
        t.Fatal("ins1 is not equal to ins2!")
    }
    ins3 := &singleton{}
    if ins1 == ins3 {
        t.Fatal("ins1 is equal to ins3!")
    }
}

/*
A quirky example:

package main

import (
	"fmt"
)

func main() {
	ins1 := &singleton{}
	ins2 := &singleton{}
	fmt.Println(ins1 == ins2)	// false

	ins3 := Singleton()
	ins4 := Singleton()
	fmt.Println(ins3 == ins4)	// true

	ins5 := new(singleton)
	ins6 := new(singleton)
	fmt.Println(ins5 == ins6)	// false

	ins7 := &singleton{}
	ins8 := &singleton{}
	fmt.Printf("%p\n", ins7)
	fmt.Printf("%p\n", ins8)
	fmt.Println(ins7 == ins8)	// true
}

type singleton struct {}

var _singleton *singleton

func Singleton() *singleton {
	_singleton = new(singleton)
	return _singleton
}

*/
