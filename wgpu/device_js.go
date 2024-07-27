//go:build js

package wgpu

import "syscall/js"

// NewDevice creates a new GPUDevice that uses the specified JavaScript
// reference of the device.
func NewDevice(jsValue js.Value) Device {
	return Device{
		jsValue: jsValue,
	}
}

// Device as described:
// https://gpuweb.github.io/gpuweb/#gpudevice
type Device struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g Device) ToJS() any {
	return g.jsValue
}

// Queue as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-queue
func (g Device) Queue() Queue {
	jsQueue := g.jsValue.Get("queue")
	return Queue{
		jsValue: jsQueue,
	}
}

// CreateCommandEncoder as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createcommandencoder
func (g Device) CreateCommandEncoder() CommandEncoder {
	jsEncoder := g.jsValue.Call("createCommandEncoder")
	return CommandEncoder{
		jsValue: jsEncoder,
	}
}

// CreateBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbuffer
func (g Device) CreateBuffer(descriptor *BufferDescriptor) (*Buffer, error) {
	jsBuffer := g.jsValue.Call("createBuffer", descriptor.ToJS())
	return &Buffer{
		jsValue: jsBuffer,
	}, nil
}

// CreateShaderModule as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createshadermodule
func (g Device) CreateShaderModule(desc ShaderModuleDescriptor) ShaderModule {
	jsShader := g.jsValue.Call("createShaderModule", desc.ToJS())
	return ShaderModule{
		jsValue: jsShader,
	}
}

// CreateRenderPipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createrenderpipeline
func (g Device) CreateRenderPipeline(descriptor RenderPipelineDescriptor) RenderPipeline {
	jsPipeline := g.jsValue.Call("createRenderPipeline", descriptor.ToJS())
	return RenderPipeline{
		jsValue: jsPipeline,
	}
}

// CreateBindGroup as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbindgroup
func (g Device) CreateBindGroup(descriptor BindGroupDescriptor) BindGroup {
	jsBindGroup := g.jsValue.Call("createBindGroup", descriptor.ToJS())
	return BindGroup{
		jsValue: jsBindGroup,
	}
}

// CreateBindGroupLayout as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbindgrouplayout
func (g Device) CreateBindGroupLayout(descriptor BindGroupLayoutDescriptor) BindGroupLayout {
	jsLayout := g.jsValue.Call("createBindGroupLayout", descriptor.ToJS())
	return BindGroupLayout{
		jsValue: jsLayout,
	}
}

// CreatePipelineLayout as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createpipelinelayout
func (g Device) CreatePipelineLayout(descriptor PipelineLayoutDescriptor) PipelineLayout {
	jsLayout := g.jsValue.Call("createPipelineLayout", descriptor.ToJS())
	return PipelineLayout{
		jsValue: jsLayout,
	}
}

// CreateComputePipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createcomputepipeline
func (g Device) CreateComputePipeline(descriptor ComputePipelineDescriptor) ComputePipeline {
	jsPipeline := g.jsValue.Call("createComputePipeline", descriptor.ToJS())
	return ComputePipeline{
		jsValue: jsPipeline,
	}
}
