//go:build js

package wgpu

import "syscall/js"

// SwapChain as described:
// https://gpuweb.github.io/gpuweb/#gpucanvascontext
// (CanvasContext is the closest equivalent to SwapChain in js)
type SwapChain struct {
	jsValue js.Value
}
