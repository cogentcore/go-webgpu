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

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g CommandEncoder) toJS() any {
	return g.jsValue
}

// BeginRenderPass as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-beginrenderpass
func (g CommandEncoder) BeginRenderPass(descriptor *RenderPassDescriptor) *RenderPassEncoder {
	jsRenderPass := g.jsValue.Call("beginRenderPass", descriptor.toJS())
	return &RenderPassEncoder{
		jsValue: jsRenderPass,
	}
}

// BeginComputePass as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-begincomputepass
func (g CommandEncoder) BeginComputePass(descriptor *ComputePassDescriptor) *ComputePassEncoder {
	params := make([]any, 1)
	params[0] = descriptor.toJS()
	jsComputePass := g.jsValue.Call("beginComputePass", params...)
	return &ComputePassEncoder{
		jsValue: jsComputePass,
	}
}

// CopyBufferToBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-copybuffertobuffer
func (g CommandEncoder) CopyBufferToBuffer(source *Buffer, sourceOffset uint64, destination *Buffer, destinationOffset uint64, size uint64) (err error) {
	g.jsValue.Call("copyBufferToBuffer", source.toJS(), sourceOffset, destination.toJS(), destinationOffset, size)
	return nil
}

// CopyBufferToTexture as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-copybuffertotexture
func (g CommandEncoder) CopyBufferToTexture(source *ImageCopyBuffer, destination *ImageCopyTexture, copySize *Extent3D) (err error) {
	g.jsValue.Call("copyBufferToTexture", source.toJS(), destination.toJS(), copySize.toJS())
	return nil
}

// CopyTextureToBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-copytexturetobuffer
func (g CommandEncoder) CopyTextureToBuffer(source *ImageCopyTexture, destination *ImageCopyBuffer, copySize *Extent3D) (err error) {
	g.jsValue.Call("copyTextureToBuffer", source.toJS(), destination.toJS(), copySize.toJS())
	return nil
}

// CopyTextureToTexture as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-copytexturetotexture
func (g CommandEncoder) CopyTextureToTexture(source *ImageCopyTexture, destination *ImageCopyTexture, copySize *Extent3D) (err error) {
	g.jsValue.Call("copyTextureToTexture", source.toJS(), destination.toJS(), copySize.toJS())
	return nil
}

// Finish as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-finish
func (g CommandEncoder) Finish(descriptor *CommandBufferDescriptor) (*CommandBuffer, error) {
	jsBuffer := g.jsValue.Call("finish", map[string]any{"label": descriptor.Label})
	return &CommandBuffer{
		jsValue: jsBuffer,
	}, nil
}

func (g CommandEncoder) Release() {} // no-op
