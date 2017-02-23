package main

import "fmt"

type myType struct {
	value int
}

func (mt myType) reset() {
	mt.value = 0
}

func (mt *myType) resetPtr() {
	mt.value = 0
}

func memberFunction() {
	mt := myType{
		value: 1,
	}
	fmt.Println("initial", mt.value)

	mt.reset()
	fmt.Println("after reset", mt.value)

	mt.resetPtr()
	fmt.Println("after resetPtr", mt.value)
}
