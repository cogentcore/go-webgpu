//go:build js

package wgpu

import "syscall/js"

// BufferDescriptor as described:
// https://gpuweb.github.io/gpuweb/#gpubufferdescriptor
type BufferDescriptor struct {
	Size  Size64
	Usage BufferUsage
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BufferDescriptor) ToJS() any {
	return map[string]any{
		"size":  g.Size.ToJS(),
		"usage": g.Usage.String(),
	}
}

// Buffer as described:
// https://gpuweb.github.io/gpuweb/#gpubuffer
type Buffer struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g Buffer) ToJS() any {
	return g.jsValue
}

// Destroy as described:
// https://gpuweb.github.io/gpuweb/#dom-gpubuffer-destroy
func (g Buffer) Destroy() {
	g.jsValue.Call("destroy")
}
