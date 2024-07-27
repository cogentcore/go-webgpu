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
	adapter := await(g.jsValue.Call("requestAdapter", pointerToJS(options)))
	if !adapter.Truthy() {
		return nil, fmt.Errorf("no WebGPU adapter avaliable")
	}
	return &Adapter{jsValue: adapter}, nil
}

type SurfaceDescriptor struct {
	Label string
}

func (g Instance) CreateSurface(descriptor *SurfaceDescriptor) *Surface {
	return &Surface{}
}

func (g Instance) GenerateReport() any { return nil } // no-op

func (g Instance) Release() {} // no-op
