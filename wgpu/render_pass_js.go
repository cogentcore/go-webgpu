//go:build js

package wgpu

import (
	"syscall/js"
)

// RenderPassColorAttachment as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpurenderpasscolorattachment
type RenderPassColorAttachment struct {
	View          TextureView
	ResolveTarget TextureView
	ClearValue    Color
	LoadOp        LoadOp
	StoreOp       StoreOp
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g RenderPassColorAttachment) ToJS() any {
	result := make(map[string]any)
	result["view"] = g.View.jsValue
	result["loadOp"] = g.LoadOp.String()
	result["storeOp"] = g.StoreOp.String()
	result["clearValue"] = g.ClearValue.ToJS()
	result["resolveTarget"] = g.ResolveTarget.ToJS()
	return result
}

// RenderPassDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpurenderpassdescriptor
type RenderPassDescriptor struct {
	ColorAttachments []RenderPassColorAttachment
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g RenderPassDescriptor) ToJS() any {
	result := make(map[string]any)
	result["colorAttachments"] = mapSlice(g.ColorAttachments, func(attachment RenderPassColorAttachment) any {
		return attachment.ToJS()
	})
	return result
}

// RenderPassEncoder as described:
// https://gpuweb.github.io/gpuweb/#gpurenderpassencoder
type RenderPassEncoder struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g RenderPassEncoder) ToJS() any {
	return g.jsValue
}

// SetPipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpurendercommandsmixin-setpipeline
func (g RenderPassEncoder) SetPipeline(pipeline RenderPipeline) {
	g.jsValue.Call("setPipeline", pipeline.ToJS())
}

// SetVertexBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpurendercommandsmixin-setvertexbuffer
func (g RenderPassEncoder) SetVertexBuffer(slot Index32, vertexBuffer Buffer, offset, size Size64) {
	params := make([]any, 4)
	params[0] = slot.ToJS()
	params[1] = vertexBuffer.ToJS()
	params[2] = offset.ToJS()
	params[3] = size.ToJS()
	g.jsValue.Call("setVertexBuffer", params...)
}

// SetBindGroup as described:
// https://gpuweb.github.io/gpuweb/#gpubindingcommandsmixin-setbindgroup
func (g RenderPassEncoder) SetBindGroup(index Index32, bindGroup BindGroup, dynamicOffsets []BufferDynamicOffset) {
	params := make([]any, 3)
	params[0] = index.ToJS()
	params[1] = bindGroup.ToJS()
	params[2] = mapSlice(dynamicOffsets, func(offset BufferDynamicOffset) any {
		return offset.ToJS()
	})
	g.jsValue.Call("setBindGroup", params...)
}

// Draw as described:
// https://gpuweb.github.io/gpuweb/#dom-gpurendercommandsmixin-draw
func (g RenderPassEncoder) Draw(vertexCount Size32, instanceCount, firstVertex, firstInstance Size32) {
	params := make([]any, 4)
	params[0] = vertexCount.ToJS()
	params[1] = instanceCount.ToJS()
	params[2] = firstVertex.ToJS()
	params[3] = firstInstance.ToJS()
	g.jsValue.Call("draw", params...)
}

// End as described:
// https://gpuweb.github.io/gpuweb/#dom-gpurenderpassencoder-end
func (g RenderPassEncoder) End() {
	g.jsValue.Call("end")
}
