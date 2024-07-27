//go:build js

package wgpu

import "syscall/js"

// CommandBuffer as described:
// https://gpuweb.github.io/gpuweb/#gpucommandbuffer
type CommandBuffer struct {
	jsValue js.Value
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g CommandBuffer) toJS() any {
	return g.jsValue
}

func (g CommandBuffer) Release() {} // no-op
