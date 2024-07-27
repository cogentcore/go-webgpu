//go:build js

package wgpu

import (
	"fmt"
	"syscall/js"
)

// Adapter as described:
// https://gpuweb.github.io/gpuweb/#gpuadapter
type Adapter struct {
	jsValue js.Value
}

func (g Adapter) RequestDevice(descriptor *DeviceDescriptor) (*Device, error) {
	device := await(g.jsValue.Call("requestDevice", descriptor.ToJS()))
	if !device.Truthy() {
		return nil, fmt.Errorf("no WebGPU device avaliable")
	}
	return &Device{jsValue: device}, nil
}

func (g *DeviceDescriptor) ToJS() js.Value {
	result := make(map[string]any)
	result["label"] = g.Label
	result["requiredFeatures"] = mapSlice(g.RequiredFeatures, func(f FeatureName) any { return f })
	// result["requiredLimits"] = // TODO(kai): convert requiredLimits to JS
	return js.ValueOf(result)
}

func (g Adapter) Release() {} // no-op

type Surface struct{} // no-op
