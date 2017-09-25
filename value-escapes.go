package main

func create() BigStruct {
	return BigStruct{}
}

func call(b BigStruct) {
	go func() {
		_ = b.value
	}()
}
