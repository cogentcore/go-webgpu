//go:build js

package wgpu

import "syscall/js"

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

func (g Buffer) GetMappedRange(offset, size uint) []byte {
	src := g.jsValue.Call("getMappedRange", offset, size)
	dst := make([]byte, src.Length())
	js.CopyBytesToGo(dst, src)
	return dst
}

func (g Buffer) Unmap() (err error) {
	g.jsValue.Call("unmap")
	return
}

func (g Buffer) Release() {} // no-op
