//go:build js

package wgpu

import (
	"log"
	"syscall/js"
)

// Instance is called GPU in js: https://gpuweb.github.io/gpuweb/#gpu-interface
type Instance struct {
	jsValue js.Value
}

func CreateInstance(descriptor *InstanceDescriptor) *Instance {
	gpu := js.Global().Get("navigator").Get("gpu")
	if !gpu.Truthy() {
		log.Println("WebGPU not supported")
		return nil
	}
	return &Instance{jsValue: gpu}
}

func (g Instance) RequestAdapter(options *RequestAdapterOptions) (*Adapter, error) {
	promise := g.jsValue.Call("requestAdapter", options.ToJS())
	then := make(chan js.Value)
	promise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		then <- args[0]
		return nil
	}))
	adapter := <-then
	if !adapter.Truthy() {
		log.Println("No WebGPU adapter available")
		return nil
	}
	return &Adapter{jsValue: then}
}

// no-op
func (g Instance) Release() {}
