//go:build js

package wgpu

import "syscall/js"

// Sampler as described:
// https://gpuweb.github.io/gpuweb/#gpusampler
type Sampler struct {
	jsValue js.Value
}
