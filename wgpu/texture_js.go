//go:build js

package wgpu

import "syscall/js"

// TextureView as described:
// https://gpuweb.github.io/gpuweb/#gputextureview
type TextureView struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g TextureView) ToJS() any {
	return g.jsValue
}

// Texture as described:
// https://gpuweb.github.io/gpuweb/#gputexture
type Texture struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g Texture) ToJS() any {
	return g.jsValue
}

// Format as described:
// https://gpuweb.github.io/gpuweb/#dom-gputexture-format
func (g Texture) Format() TextureFormat {
	jsFormat := g.jsValue.Get("format")
	return TextureFormat(jsFormat.Int()) // TODO(kai): need to set from string
}

// CreateView as described:
// https://gpuweb.github.io/gpuweb/#dom-gputexture-createview
func (g Texture) CreateView() TextureView {
	jsView := g.jsValue.Call("createView")
	return TextureView{
		jsValue: jsView,
	}
}
