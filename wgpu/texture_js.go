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

func (g TextureView) Release() {} // no-op

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

// GetFormat as described:
// https://gpuweb.github.io/gpuweb/#dom-gputexture-format
func (g Texture) GetFormat() TextureFormat {
	jsFormat := g.jsValue.Get("format")
	return TextureFormat(jsFormat.Int()) // TODO(kai): need to set from string
}

// CreateView as described:
// https://gpuweb.github.io/gpuweb/#dom-gputexture-createview
func (g Texture) CreateView(descriptor *TextureViewDescriptor) (*TextureView, error) {
	jsView := g.jsValue.Call("createView", descriptor.ToJS())
	return &TextureView{
		jsValue: jsView,
	}, nil
}

func (g *TextureViewDescriptor) ToJS() js.Value {
	return js.ValueOf(map[string]any{
		"label":           g.Label,
		"format":          g.Format.String(),
		"dimension":       g.Dimension.String(),
		"baseMipLevel":    g.BaseMipLevel,
		"mipLevelCount":   g.MipLevelCount,
		"baseArrayLayer":  g.BaseArrayLayer,
		"arrayLayerCount": g.ArrayLayerCount,
		"aspect":          g.Aspect.String(),
	})

}

func (g Texture) Release() {} // no-op
