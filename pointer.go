package main

func create() *BigStruct {
	return &BigStruct{}
}

func call(b *BigStruct) {
	_ = b.value
}
