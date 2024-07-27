//go:build js

package wgpu

import (
	"syscall/js"
)

// CommandEncoder as described:
// https://gpuweb.github.io/gpuweb/#gpucommandencoder
type CommandEncoder struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g CommandEncoder) ToJS() any {
	return g.jsValue
}

// BeginRenderPass as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-beginrenderpass
func (g CommandEncoder) BeginRenderPass(descriptor RenderPassDescriptor) RenderPassEncoder {
	jsRenderPass := g.jsValue.Call("beginRenderPass", descriptor.ToJS())
	return RenderPassEncoder{
		jsValue: jsRenderPass,
	}
}

// BeginComputePass as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-begincomputepass
func (g CommandEncoder) BeginComputePass(descriptor ComputePassDescriptor) ComputePassEncoder {
	params := make([]any, 1)
	params[0] = descriptor.ToJS()
	jsComputePass := g.jsValue.Call("beginComputePass", params...)
	return ComputePassEncoder{
		jsValue: jsComputePass,
	}
}

// Finish as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-finish
func (g CommandEncoder) Finish() CommandBuffer {
	jsBuffer := g.jsValue.Call("finish")
	return CommandBuffer{
		jsValue: jsBuffer,
	}
}
