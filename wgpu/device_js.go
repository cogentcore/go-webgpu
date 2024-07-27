//go:build js

package wgpu

import (
	"syscall/js"
)

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

func (g Device) toJS() any {
	return g.jsValue
}

// Queue as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-queue
func (g Device) GetQueue() *Queue {
	jsQueue := g.jsValue.Get("queue")
	return &Queue{
		jsValue: jsQueue,
	}
}

// CreateCommandEncoder as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createcommandencoder
func (g Device) CreateCommandEncoder(descriptor *CommandEncoderDescriptor) (*CommandEncoder, error) {
	jsEncoder := g.jsValue.Call("createCommandEncoder", toJS(descriptor))
	return &CommandEncoder{
		jsValue: jsEncoder,
	}, nil
}

// CreateBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbuffer
func (g Device) CreateBuffer(descriptor *BufferDescriptor) (*Buffer, error) {
	jsBuffer := g.jsValue.Call("createBuffer", toJS(descriptor))
	return &Buffer{
		jsValue: jsBuffer,
	}, nil
}

// CreateShaderModule as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createshadermodule
func (g Device) CreateShaderModule(desc *ShaderModuleDescriptor) (*ShaderModule, error) {
	jsShader := g.jsValue.Call("createShaderModule", toJS(desc))
	return &ShaderModule{
		jsValue: jsShader,
	}, nil
}

// CreateRenderPipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createrenderpipeline
func (g Device) CreateRenderPipeline(descriptor *RenderPipelineDescriptor) (*RenderPipeline, error) {
	jsPipeline := g.jsValue.Call("createRenderPipeline", toJS(descriptor))
	return &RenderPipeline{
		jsValue: jsPipeline,
	}, nil
}

// CreateBindGroup as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbindgroup
func (g Device) CreateBindGroup(descriptor *BindGroupDescriptor) (*BindGroup, error) {
	jsBindGroup := g.jsValue.Call("createBindGroup", toJS(descriptor))
	return &BindGroup{
		jsValue: jsBindGroup,
	}, nil
}

// CreateBindGroupLayout as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbindgrouplayout
func (g Device) CreateBindGroupLayout(descriptor *BindGroupLayoutDescriptor) (*BindGroupLayout, error) {
	jsLayout := g.jsValue.Call("createBindGroupLayout", toJS(descriptor))
	return &BindGroupLayout{
		jsValue: jsLayout,
	}, nil
}

// CreatePipelineLayout as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createpipelinelayout
func (g Device) CreatePipelineLayout(descriptor *PipelineLayoutDescriptor) (*PipelineLayout, error) {
	jsLayout := g.jsValue.Call("createPipelineLayout", toJS(descriptor))
	return &PipelineLayout{
		jsValue: jsLayout,
	}, nil
}

// CreateComputePipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createcomputepipeline
func (g Device) CreateComputePipeline(descriptor *ComputePipelineDescriptor) (*ComputePipeline, error) {
	jsPipeline := g.jsValue.Call("createComputePipeline", toJS(descriptor))
	return &ComputePipeline{
		jsValue: jsPipeline,
	}, nil
}

// CreateTexture as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createtexture
func (g Device) CreateTexture(descriptor *TextureDescriptor) (*Texture, error) {
	jsTexture := g.jsValue.Call("createTexture", toJS(descriptor))
	return &Texture{
		jsValue: jsTexture,
	}, nil
}

func (g Device) Poll(wait bool, wrappedSubmissionIndex *WrappedSubmissionIndex) (queueEmpty bool) {
	return false // no-op
}

func (g Device) Release() {} // no-op
