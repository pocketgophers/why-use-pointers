package main_test

import "testing"

type BigStruct struct {
	value int
	stuff [1000000]int
}

// calling this will make a copy of b
func value(b BigStruct) {
	_ = b.value
}

// calling this will only make a copy of the pointer to b
func valuePointer(b *BigStruct) {
	_ = b.value
}

func BenchmarkCopy(B *testing.B) {
	b := BigStruct{}

	for i := 0; i < B.N; i++ {
		value(b)
	}
}

func BenchmarkPointer(B *testing.B) {
	b := &BigStruct{}

	for i := 0; i < B.N; i++ {
		valuePointer(b)
	}
}

func valueEscapes(b BigStruct) {
	// running in a goroutine to convince the escape analysis that
	// `b` escapes
	go func() {
		_ = b.value
	}()
}

func valuePointerEscapes(b *BigStruct) {
	go func() {
		_ = b.value
	}()
}

func BenchmarkCopyEscapes(B *testing.B) {
	b := BigStruct{}

	for i := 0; i < B.N; i++ {
		valueEscapes(b)
	}
}

func BenchmarkPointerEscapes(B *testing.B) {
	b := &BigStruct{}

	for i := 0; i < B.N; i++ {
		valuePointerEscapes(b)
	}
}
