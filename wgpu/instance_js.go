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
	adapter := g.jsValue.Call("requestAdapter", options.ToJS())
	if !adapter.Truthy() {
		return nil
	}
	return &Adapter{jsValue: adapter}
}

// no-op
func (g Instance) Release() {}
