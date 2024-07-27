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
func (g CommandEncoder) BeginRenderPass(descriptor *RenderPassDescriptor) *RenderPassEncoder {
	jsRenderPass := g.jsValue.Call("beginRenderPass", descriptor.ToJS())
	return &RenderPassEncoder{
		jsValue: jsRenderPass,
	}
}

// BeginComputePass as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-begincomputepass
func (g CommandEncoder) BeginComputePass(descriptor *ComputePassDescriptor) *ComputePassEncoder {
	params := make([]any, 1)
	params[0] = descriptor.ToJS()
	jsComputePass := g.jsValue.Call("beginComputePass", params...)
	return &ComputePassEncoder{
		jsValue: jsComputePass,
	}
}

// CopyBufferToBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-copybuffertobuffer
func (g CommandEncoder) CopyBufferToBuffer(source *Buffer, sourceOffset uint64, destination *Buffer, destinationOffset uint64, size uint64) (err error) {
	g.jsValue.Call("copyBufferToBuffer", source.ToJS(), sourceOffset, destination.ToJS(), destinationOffset, size)
	return nil
}

// CopyBufferToTexture as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-copybuffertotexture
func (g CommandEncoder) CopyBufferToTexture(source *ImageCopyBuffer, destination *ImageCopyTexture, copySize *Extent3D) (err error) {
	g.jsValue.Call("copyBufferToTexture", source.ToJS(), destination.ToJS(), copySize.ToJS())
	return nil
}

// CopyTextureToBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-copytexturetobuffer
func (g CommandEncoder) CopyTextureToBuffer(source *ImageCopyTexture, destination *ImageCopyBuffer, copySize *Extent3D) (err error) {
	g.jsValue.Call("copyTextureToBuffer", source.ToJS(), destination.ToJS(), copySize.ToJS())
	return nil
}

// CopyTextureToTexture as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-copytexturetotexture
func (g CommandEncoder) CopyTextureToTexture(source *ImageCopyTexture, destination *ImageCopyTexture, copySize *Extent3D) (err error) {
	g.jsValue.Call("copyTextureToTexture", source.ToJS(), destination.ToJS(), copySize.ToJS())
	return nil
}

func (g *ImageCopyBuffer) ToJS() any {
	return map[string]any{
		"layout": g.Layout.ToJS(),
		"buffer": g.Buffer.ToJS(),
	}
}

func (g *ImageCopyTexture) ToJS() any {
	return map[string]any{
		"texture":  g.Texture.ToJS(),
		"mipLevel": g.MipLevel,
		"origin":   g.Origin.ToJS(),
		"aspect":   g.Aspect.String(),
	}
}

func (g *TextureDataLayout) ToJS() any {
	return map[string]any{
		"offset":       g.Offset,
		"bytesPerRow":  g.BytesPerRow,
		"rowsPerImage": g.RowsPerImage,
	}
}

func (g *Origin3D) ToJS() any {
	return []any{g.X, g.Y, g.Z}
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
