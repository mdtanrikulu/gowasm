package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	msg := "Hello, WebAssembly lol!"
	fmt.Println(msg)
	document := js.Global().Get("document")
	newSpan := document.Call("createElement", "span")
	newSpan.Set("innerText", msg)
	document.Get("body").Call("appendChild", newSpan)
}
