//go:build js

package wgpu

import (
	"fmt"
	"log"
	"syscall/js"
)

// Instance as described:
// https://gpuweb.github.io/gpuweb/#gpu-interface
// (Instance is called GPU in js)
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
		return nil, fmt.Errorf("no WebGPU adapter avaliable")
	}
	return &Adapter{jsValue: adapter}, nil
}

func (g Instance) Release() {} // no-op

func (g RequestAdapterOptions) ToJS() js.Value {
	result := make(map[string]any)
	result["powerPreference"] = g.PowerPreference.String()
	result["forceFallbackAdapter"] = g.ForceFallbackAdapter
	return js.ValueOf(result)
}
