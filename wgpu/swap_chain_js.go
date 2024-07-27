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
	texture := &Texture{jsValue: g.jsValue.Call("getCurrentTexture")}
	return texture.CreateView(&TextureViewDescriptor{}) // TODO(kai): set attributes
}

func (g SwapChain) Present() {} // no-op

func (g SwapChain) Release() {} // no-op
