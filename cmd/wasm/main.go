package main

import (
	"fmt"
	"syscall/js"
)

func superMultiply(a int, b int) int {
	return a * b
}

// gets 2 values from js and returns the result
func superMultiplyWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return superMultiply(args[0].Int(), args[1].Int())
	})
}

// gets 1 object { a, b } from js and returns the result
func superMultiplyObjectWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		type MyObject struct {
			a int
			b int
		}
		myObj := MyObject{
			a: args[0].Get("a").Int(),
			b: args[0].Get("b").Int(),
		}
		return superMultiply(myObj.a, myObj.b)
	})
}

// gets 1 array [ a, b ] from js and returns the result
func superMultiplyArrayWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		length := args[0].Get("length").Int()
		input := make([]int, length)
		for i := 0; i < length; i++ {
			input[i] = args[0].Index(i).Int()
		}
		return superMultiply(input[0], input[1])
	})
}

// gets 1 array of objects [ {value: a}, {value: b} ] from js and returns the result
func superMultiplyArrayOfObjectsWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		type MyObject struct {
			value int
		}
		length := args[0].Get("length").Int()
		input := make([]MyObject, length)
		for i := 0; i < length; i++ {
			//input[i] = args[0].Index(i).Int()
			input[i] = MyObject{
				value: args[0].Index(i).Get("value").Int(),
			}
		}
		return superMultiply(input[0].value, input[1].value)
	})
}

func main() {
	fmt.Println("Super multiplier in assembly")
	fmt.Println("- superMultiplySimple")
	fmt.Println("- superMultiplyObject")
	fmt.Println("- superMultiplyArray")
	fmt.Println("- superMultiplyArrayOfObjects")
	js.Global().Set("superMultiplySimple", superMultiplyWrapper())
	js.Global().Set("superMultiplyObject", superMultiplyObjectWrapper())
	js.Global().Set("superMultiplyArray", superMultiplyArrayWrapper())
	js.Global().Set("superMultiplyArrayOfObjects", superMultiplyArrayOfObjectsWrapper())
	<-make(chan bool)
}
