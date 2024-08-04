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
	device := await(g.jsValue.Call("requestDevice", pointerToJS(descriptor)))
	if !device.Truthy() {
		return nil, fmt.Errorf("no WebGPU device avaliable")
	}
	return &Device{jsValue: device}, nil
}

func (g Adapter) GetProperties() AdapterProperties {
	return AdapterProperties{} // TODO(kai): implement?
}

func (g Adapter) GetLimits() SupportedLimits {
	return SupportedLimits{limitsFromJS(g.jsValue.Get("limits"))}
}

func (g Adapter) Release() {} // no-op
