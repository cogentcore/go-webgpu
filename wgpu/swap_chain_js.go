//go:build js

package wgpu

import "syscall/js"

// SwapChain as described:
// https://gpuweb.github.io/gpuweb/#gpucanvascontext
// (CanvasContext is the closest equivalent to SwapChain in js)
type SwapChain struct {
	jsValue js.Value
}

func (g SwapChain) GetCurrentTextureView() (*TextureView, error) {
	texture := g.jsValue.Call("getCurrentTexture")
	// We can just use the properties of the texture as the descriptor.
	desc := TextureViewDescriptor{Format: TextureFormatBGRA8UnormSrgb, Dimension: TextureViewDimension2D, Aspect: TextureAspectAll}
	return &TextureView{jsValue: texture.Call("createView", texture, desc.toJS())}, nil
}

func (g SwapChain) Present() {} // no-op

func (g SwapChain) Release() {} // no-op
