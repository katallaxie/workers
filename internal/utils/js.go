package utils

import (
	"fmt"
	"syscall/js"
)

var (
	Global              = js.Global()
	ObjectClass         = Global.Get("Object")
	ArrayClass          = Global.Get("Array")
	FunctionClass       = Global.Get("Function")
	NumberClass         = Global.Get("Number")
	PromiseClass        = Global.Get("Promise")
	RequestClass        = Global.Get("Request")
	ResponseClass       = Global.Get("Response")
	HeadersClass        = Global.Get("Headers")
	Uint8ArrayClass     = Global.Get("Uint8Array")
	ErrorClass          = Global.Get("Error")
	ReadableStreamClass = Global.Get("ReadableStream")
	DateClass           = Global.Get("Date")
	Undefined           = js.Undefined()
	Null                = js.ValueOf(nil)
)

// NewObject returns a new JavaScript object.
func NewObject() js.Value {
	return ObjectClass.New()
}

// NewUint8Array returns a new Uint8Array.
func NewUint8Array(size int) js.Value {
	return Uint8ArrayClass.New(size)
}

// NewPromise returns a new Promise.
func NewPromise(fn js.Func) js.Value {
	return PromiseClass.New(fn)
}

// ArrayFrom returns a new Array from a given iterable.
func ArrayFrom(iterable js.Value) js.Value {
	return ArrayClass.Call("from", iterable)
}

// Await ...
func Await(promise js.Value) (js.Value, error) {
	result := make(chan js.Value)
	err := make(chan error)

	var then, catch js.Func
	then = js.FuncOf(func(_ js.Value, args []js.Value) any {
		defer then.Release()
		result := args[0]
		result <- result
		return js.Undefined()
	})
	catch = js.FuncOf(func(_ js.Value, args []js.Value) any {
		defer catch.Release()
		result := args[0]
		err <- fmt.Errorf("failed on promise: %s", result.Call("toString").String())
		return js.Undefined()
	})

	promise.Call("then", then).Call("catch", catch)

	select {
	case result := <-result:
		return result, nil
	case err := <-err:
		return js.Value{}, err
	}
}
