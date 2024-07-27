//go:build js

package wgpu

import (
	"syscall/js"
)

// Queue as described:
// https://gpuweb.github.io/gpuweb/#gpuqueue
type Queue struct {
	jsValue js.Value
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g Queue) toJS() any {
	return g.jsValue
}

// Submit as described:
// https://gpuweb.github.io/gpuweb/#dom-gpuqueue-submit
func (g Queue) Submit(commandBuffers ...*CommandBuffer) {
	jsSequence := mapSlice(commandBuffers, func(buffer *CommandBuffer) any {
		return buffer.toJS()
	})
	g.jsValue.Call("submit", jsSequence)
}

// WriteBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpuqueue-writebuffer
func (g Queue) WriteBuffer(buffer Buffer, offset uint64, data []byte) {
	dataSize := stageBufferData(data)
	g.jsValue.Call("writeBuffer", buffer.jsValue, offset, uint8Array, uint64(0), dataSize)
}

func (g Queue) Release() {} // no-op
