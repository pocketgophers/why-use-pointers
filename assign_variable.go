package main

import "fmt"

func reset(value int) {
	value = 0
}

func resetPtr(value *int) {
	*value = 0
}

func assignVariable() {
	value := 1
	fmt.Println("initial", value)

	reset(value)
	fmt.Println("after reset", value)

	resetPtr(&value)
	fmt.Println("after resetPtr", value)
}
