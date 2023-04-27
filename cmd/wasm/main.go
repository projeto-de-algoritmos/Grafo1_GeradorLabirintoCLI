package main

import (
	"fmt"
	"syscall/js"
)

func LeftButom() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func UpButom() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func DowmButom() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func RightButom() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func RightClick() {
	doc := js.Global().Get("document")
	// Get the button element by its ID
	button := doc.Call("getElementById", "right")

	// Add an event listener to the button for the "click" event
	button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// This function will be called when the button is clicked
		println("Button right")
		return nil
	}))
}

func LeftClick() {
	doc := js.Global().Get("document")
	// Get the button element by its ID
	button := doc.Call("getElementById", "left")

	// Add an event listener to the button for the "click" event
	button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// This function will be called when the button is clicked
		println("Button clicked! 23")
		return nil
	}))
}

func UpClick() {
	doc := js.Global().Get("document")
	// Get the button element by its ID
	button := doc.Call("getElementById", "up")

	// Add an event listener to the button for the "click" event
	button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// This function will be called when the button is clicked
		println("Button up")
		return nil
	}))
}

func DownClick() {
	doc := js.Global().Get("document")
	// Get the button element by its ID
	button := doc.Call("getElementById", "down")

	// Add an event listener to the button for the "click" event
	button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// This function will be called when the button is clicked
		println("Button dowm")
		return nil
	}))
}

func main() {

	ch := make(chan struct{}, 0)                // mantem rodando o programa infinitamente
	fmt.Printf("Hello Web Assembly from Go!\n") // printa na console

	js.Global().Set("left", LeftButom())
	js.Global().Set("down", DowmButom())
	js.Global().Set("right", RightButom())
	js.Global().Set("up", UpButom())

	RightClick()
	LeftClick()
	UpClick()
	DownClick()
	Map()
	<-ch
}
