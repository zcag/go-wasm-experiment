package main

import (
	"syscall/js"
)

func calc(this js.Value, args []js.Value) any {
	a := args[0].Int()
	b := args[1].Int()
	return js.ValueOf(a*a + b*b)
}

func main() {
	js.Global().Set("calc", js.FuncOf(calc))
	select {} // keep Go runtime alive
}
